package routes

import (
	"rest-api/user"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoute(r *gin.Engine, handler *user.UserHandler) {
	r.GET("/users", handler.GetUsers)

	r.GET("/users/:id", handler.GetUserByID)
}
