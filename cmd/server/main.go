package main

import (
	"log"
	"rest-api/hello"
	"rest-api/routes"
	"rest-api/user"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Register routes from other packages
	routes.RegisterHelloRoute(r, InitializeHelloHandler())

	routes.RegisterUserRoute(r, InitializeUserHandler())

	if err := r.Run(":3600"); err != nil {
		log.Fatal(err)
	}
}

func InitializeHelloHandler() *hello.HelloHandler {
	return hello.NewHelloHandler()
}

func InitializeUserHandler() *user.UserHandler {
	userService := user.NewUserService()
	return user.NewUserHandler(userService)
}
