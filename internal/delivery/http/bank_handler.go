package http

import (
	"demo/bank-linking-listener/internal/delivery/http/http_dto"
	"demo/bank-linking-listener/internal/service"
	"demo/bank-linking-listener/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BankHandler struct {
	bankService service.BankService
}

func NewBankHandler(bankSerice service.BankService) *BankHandler {
	return &BankHandler{bankService: bankSerice}
}

func (h *BankHandler) CheckHealth(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *BankHandler) GetBankListCurrentUser(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *BankHandler) GetBankListByUserID(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *BankHandler) CreateBank(c *gin.Context) {
	var req http_dto.BankCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseWithMessage(
			utils.ResponseStatusFail, err.Error()))
		return
	}

	bank, rerr := h.bankService.CreateBank(c.Request.Context(), *req.ToEntity())
	if rerr != nil {
		c.JSON(utils.GetStatusCode(rerr), utils.ResponseWithMessage(
			utils.ResponseStatusFail, rerr.Message()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseWithData(
		utils.ResponseStatusSuccess, map[string]interface{}{
			"bank": bank,
		},
	))
}
