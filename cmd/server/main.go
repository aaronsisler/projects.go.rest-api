package main

import (
	"log"
	"rest-api/routes"
	"rest-api/user"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Register routes from other packages
	routes.RegisterHelloRoute(r)

	// Instantiate service and handler
	userHandler := InitializeUserHandler()
	routes.RegisterUserRoute(r, userHandler)

	if err := r.Run(":3600"); err != nil {
		log.Fatal(err)
	}
}

func InitializeUserHandler() *user.UserHandler {
	userService := user.NewUserService()
	return user.NewUserHandler(userService)
}
