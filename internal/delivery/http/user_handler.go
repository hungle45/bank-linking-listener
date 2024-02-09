package http

import (
	"demo/bank-linking-listener/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) CheckHealth(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
