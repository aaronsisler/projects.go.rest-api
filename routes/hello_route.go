package routes

import (
	"rest-api/hello"

	"github.com/gin-gonic/gin"
)

// RegisterHelloRoute adds the /hello route to the given router group
func RegisterHelloRoute(r *gin.Engine, handler *hello.HelloHandler) {
	r.GET("/hello", handler.GetHello)
}
