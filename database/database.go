package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Query struct {
	db *gorm.DB
}

func NewQuery() *Query {
	db, err := gorm.Open("sqlite3", "dev.db")

	if err != nil {
		panic("Error : can't connect to database dev.db")
	}

	db.AutoMigrate(&User{})

	query := new(Query)
	query.db = db

	return query
}

func (this *Query) Destroy() {
	this.db.Close()
}

func (this *Query) Create(username, email string) User {
	user := User{Username: username, Email: email}

	this.db.Create(&user)
	return user
}

func (this *Query) GetByID(id int) User {
	var user User

	this.db.First(&user, id)
	return user
}

func (this *Query) GetAll() []User {
	var users []User

	this.db.Find(&users)
	return users
}

func (this *Query) Update(id int, username, email string) User {
	var user User
	updateMap := map[string]interface{}{"username": username, "email": email}

	this.db.Model(&user).Where("id = ?", id).Update(updateMap)
	user.ID = uint(id)
	return user
}

func (this *Query) Delete(id int) {
	var user User

	this.db.First(&user, id)
	this.db.Delete(&user)
}
