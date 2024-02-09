package http

import (
	"demo/bank-linking-listener/internal/delivery/http/http_dto"
	"demo/bank-linking-listener/internal/service"
	"demo/bank-linking-listener/pkg/utils"
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

func (h *UserHandler) SignUp(c *gin.Context) {
	var req http_dto.UserSignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseWithMessage(
			utils.ResponseStatusFail, err.Error()))
		return
	}

	if rerr := h.userService.SignUp(c.Request.Context(), *req.ToEntity()); rerr != nil {
		c.JSON(utils.GetStatusCode(rerr), utils.ResponseWithMessage(
			utils.ResponseStatusFail, rerr.Message()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseWithMessage(
		utils.ResponseStatusSuccess, "account has been created"))
}

func (h *UserHandler) SignIn(c *gin.Context) {
	var req http_dto.UserSignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseWithMessage(
			utils.ResponseStatusFail, err.Error()))
		return
	}

	token, rerr := h.userService.SignIn(c.Request.Context(), *req.ToEntity())
	if rerr != nil {
		c.JSON(utils.GetStatusCode(rerr), utils.ResponseWithMessage(
			utils.ResponseStatusFail, rerr.Message()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseWithData(
		utils.ResponseStatusSuccess, map[string]interface{}{"token": token}))
}
