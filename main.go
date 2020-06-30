package main

import (
	"./database"
)

func main() {
	db := database.Init()
	defer database.Destroy(db)
}
