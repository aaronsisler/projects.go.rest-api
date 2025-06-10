package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterHelloRoute adds the /hello route to the given router group
func RegisterHelloRoute(r *gin.Engine) {
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})
}
