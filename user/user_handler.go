package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *UserService
}

func NewUserHandler(userService *UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user := h.UserService.GetUser(id)
	c.JSON(http.StatusOK, gin.H{"user": user})
}
