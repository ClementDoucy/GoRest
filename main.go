package main

import (
	"fmt"

	"./app"
)

func main() {
	a := app.App{}

	a.Init()

	fmt.Println("Running on http://0.0.0.0:8080")

	a.Run(":8080")
}
