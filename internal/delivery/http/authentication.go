package http

import (
	"demo/bank-linking-listener/internal/delivery/http/http_dto"
	"demo/bank-linking-listener/pkg/errorx"
	"demo/bank-linking-listener/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller *Controller) CreateUserAccount(c *gin.Context) {
	var req http_dto.UserSignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseWithMessage(
			utils.ResponseStatusFail, err.Error()))
		return
	}

	if err := controller.userService.CreateUserAccount(c.Request.Context(), *req.ToEntity()); err != nil {
		c.JSON(errorx.GetHTTPCode(err), utils.ResponseWithMessage(
			utils.ResponseStatusFail, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseWithMessage(
		utils.ResponseStatusSuccess, "account has been created"))
}

func (controller *Controller) CreateCustomerAccount(c *gin.Context) {
	var req http_dto.UserSignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseWithMessage(
			utils.ResponseStatusFail, err.Error()))
		return
	}

	if err := controller.userService.CreateCustomerAccount(c.Request.Context(), *req.ToEntity()); err != nil {
		c.JSON(errorx.GetHTTPCode(err), utils.ResponseWithMessage(
			utils.ResponseStatusFail, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseWithMessage(
		utils.ResponseStatusSuccess, "account has been created"))
}

func (controller *Controller) SignIn(c *gin.Context) {
	var req http_dto.UserSignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseWithMessage(
			utils.ResponseStatusFail, err.Error()))
		return
	}

	token, err := controller.userService.SignIn(c.Request.Context(), *req.ToEntity())
	if err != nil {
		c.JSON(errorx.GetHTTPCode(err), utils.ResponseWithMessage(
			utils.ResponseStatusFail, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseWithData(
		utils.ResponseStatusSuccess, map[string]interface{}{"token": token}))
}
