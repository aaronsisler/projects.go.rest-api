package main

import (
	"log"
	"rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	userHandler := InitializeUserHandler() // Wire-injected

	r.GET("/users/:id", userHandler.GetUser)

	// Register routes from other packages
	routes.RegisterHelloRoute(r)

	if err := r.Run(":3600"); err != nil {
		log.Fatal(err)
	}
}
