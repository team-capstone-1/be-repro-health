package main

import (
	"capstone-project/database"
	m "capstone-project/middleware"
	"capstone-project/route"
	"capstone-project/config"
)

func main() {
	config.Init()
	database.Init()
	e := route.New()
	//implement middleware logger
	m.LogMiddlewares(e)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}
