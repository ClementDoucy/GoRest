package main

import (
	"./database"
)

func main() {
	query := database.NewQuery()
	defer query.Destroy()
}
