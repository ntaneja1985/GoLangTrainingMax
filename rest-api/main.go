package main

import (
	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	//Initialize the database
	db.InitDB()
	//Behind the scenes configures an HTTP server for us
	//with Logger and Recovery middleware already attached
	//Returns a handle(pointer) to the engine or the server
	server := gin.Default()

	//Register the Routes
	routes.RegisterRoutes(server)

	//We can use the handle to run the server at port 8080
	server.Run(":8080") //localhost:8080

}
