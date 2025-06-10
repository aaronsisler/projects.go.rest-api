package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterHelloRoute adds the /hello route to the given router group
func RegisterUserRoute(r *gin.Engine) {
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusNoContent, nil)
	})

	r.GET("/users/:id", getUserByID)
}

// getUserByID handles GET /users/:id
func getUserByID(c *gin.Context) {
	id := c.Param("id") // get the path parameter

	// You could validate or lookup a user here. For now, just echo it.
	c.JSON(http.StatusOK, gin.H{
		"user_id": id,
		"message": "User fetched successfully",
	})
}
