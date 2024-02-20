package http

import (
	"demo/bank-linking-listener/internal/delivery/http/http_dto"
	"demo/bank-linking-listener/pkg/errorx"
	"demo/bank-linking-listener/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller *Controller) CreateBank(c *gin.Context) {
	var req http_dto.BankCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseWithMessage(
			utils.ResponseStatusFail, err.Error()))
		return
	}

	bank, err := controller.bankService.CreateBank(c.Request.Context(), *req.ToEntity())
	if err != nil {
		c.JSON(errorx.GetHTTPCode(err), utils.ResponseWithMessage(
			utils.ResponseStatusFail, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseWithData(
		utils.ResponseStatusSuccess, map[string]interface{}{
			"bank": bank,
		},
	))
}
