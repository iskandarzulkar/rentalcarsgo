package main

import (
	"IskandarGo/db"
	"IskandarGo/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":8080"))
}
