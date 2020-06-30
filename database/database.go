package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	ID       uint `gorm:"primary_key"`
	Username string
	Email    string
}

func Init() *gorm.DB {
	db, err := gorm.Open("sqlite3", "dev.db")

	if err != nil {
		panic("Error : can't connect to database")
	}

	db.AutoMigrate(&User{})

	return db
}

func Destroy(db *gorm.DB) {
	db.Close()
}

func Create(db *gorm.DB, username, email string) User {
	user := User{Username: username, Email: email}

	db.Create(&user)
	return user
}

func GetByID(db *gorm.DB, id int) User {
	var user User

	db.First(&user, id)
	return user
}

func GetAll(db *gorm.DB) []User {
	var users []User

	db.Find(&users)
	return users
}

func Update(db *gorm.DB, id int, username, email string) User {
	var user User
	updateMap := map[string]interface{}{"username": username, "email": email}

	db.Model(&user).Where("id = ?", id).Update(updateMap)
	return user
}

func Delete(db *gorm.DB, id int) {
	var user User

	db.First(&user, id)
	db.Delete(&user)
}
