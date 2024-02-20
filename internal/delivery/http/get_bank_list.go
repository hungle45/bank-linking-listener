package http

import (
	"demo/bank-linking-listener/pkg/errorx"
	"demo/bank-linking-listener/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (controller *Controller) GetBankListCurrentUser(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	banks, err := controller.bankService.GetBankListByUserID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(errorx.GetHTTPCode(err), utils.ResponseWithMessage(
			utils.ResponseStatusFail, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseWithData(
		utils.ResponseStatusSuccess, map[string]interface{}{
			"banks": banks,
		},
	))
}

func (controller *Controller) GetBankListByUserID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil || userID <= 0 {
		c.JSON(http.StatusBadRequest, utils.ResponseWithMessage(
			utils.ResponseStatusFail, "invalid user id"))
	}

	banks, err := controller.bankService.GetBankListByUserID(c.Request.Context(), uint(userID))
	if err != nil {
		c.JSON(errorx.GetHTTPCode(err), utils.ResponseWithMessage(
			utils.ResponseStatusFail, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseWithData(
		utils.ResponseStatusSuccess, map[string]interface{}{
			"user_id": userID,
			"banks":   banks,
		},
	))
}
