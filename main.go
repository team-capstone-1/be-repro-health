package main

import (
	"capstone-project/database"
	m "capstone-project/middleware"
	"capstone-project/route"
)

func main() {
	database.InitTest()
	e := route.New()
	//implement middleware logger
	m.LogMiddlewares(e)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}
