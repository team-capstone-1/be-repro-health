package main

import (
	m "capstone-project/middleware"
	"capstone-project/route"
	"capstone-project/database"
)

func main() {
	database.Init()
	e := route.New()
	//implement middleware logger
	m.LogMiddlewares(e)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}
