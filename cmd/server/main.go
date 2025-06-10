package main

import (
	"rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Register routes from other packages
	routes.RegisterHelloRoute(r)

	r.Run(":3600") // starts the server
}
