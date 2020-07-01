package middlewares

import (
	"net/http"
	"strconv"

	"../response"

	"github.com/gorilla/mux"
)

func CheckID(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)

			id, err := strconv.Atoi(vars["id"])

			if err != nil || id < 0 {
				response.JSON(w, http.StatusBadRequest, "Invalid ID")
			} else {
				next.ServeHTTP(w, r)
			}
			return
		})
}
