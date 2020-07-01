package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"../database"
	"../middlewares"
	"../response"
	"github.com/gorilla/mux"
)

type App struct {
	router *mux.Router
	query  *database.Query
}

func (this *App) Init() {
	this.query = database.NewQuery()
	this.router = mux.NewRouter()

	this.router.HandleFunc("/", this.GetAllUser).Methods("GET")
	this.router.HandleFunc("/{id}", middlewares.CheckID(this.GetUser)).Methods("GET")
	this.router.HandleFunc("/", this.CreateUser).Methods("POST")
	this.router.HandleFunc("/{id}", middlewares.CheckID(this.UpdateUser)).Methods("PUT")
	this.router.HandleFunc("/{id}", middlewares.CheckID(this.DeleteUser)).Methods("DELETE")
}

func (this *App) Run(port string) {
	log.Fatal(http.ListenAndServe(port, this.router))
}

func (this *App) GetAllUser(w http.ResponseWriter, r *http.Request) {
	users := this.query.GetAll()

	response.JSON(w, http.StatusOK, users)
}

func (this *App) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	user := this.query.GetByID(id)

	if user.ID != 0 {
		response.JSON(w, http.StatusOK, user)
	} else {
		response.JSON(w, http.StatusNotFound, fmt.Sprintf("User %d doesn't exist.", id))
	}
}

func (this *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var Data struct {
		Username string
		Email    string
	}

	if json.NewDecoder(r.Body).Decode(&Data) != nil {
		response.JSON(w, http.StatusBadRequest, "Please send JSON data")
	} else if Data.Username == "" || Data.Email == "" {
		response.JSON(w, http.StatusBadRequest, "Please set username and email values")
	} else {
		user := this.query.Create(Data.Username, Data.Email)
		response.JSON(w, http.StatusCreated, user)
	}
}

func (this *App) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var Data struct {
		Username string
		Email    string
	}

	if json.NewDecoder(r.Body).Decode(&Data) != nil {
		response.JSON(w, http.StatusBadRequest, "Please send JSON data")
	} else if Data.Username == "" || Data.Email == "" {
		response.JSON(w, http.StatusBadRequest, "Please set username and email values")
	} else {
		user := this.query.GetByID(id)
		if user.ID == 0 {
			response.JSON(w, http.StatusNotFound, fmt.Sprintf("User %d doesn't exist.", id))
		} else {
			user = this.query.Update(id, Data.Username, Data.Email)
			response.JSON(w, http.StatusOK, user)
		}
	}
}

func (this *App) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	user := this.query.GetByID(id)
	if user.ID == 0 {
		response.JSON(w, http.StatusNotFound, fmt.Sprintf("User %d doesn't exist.", id))
	} else {
		this.query.Delete(id)
		response.JSON(w, http.StatusOK, fmt.Sprintf("User %d deleted.", id))
	}
}
