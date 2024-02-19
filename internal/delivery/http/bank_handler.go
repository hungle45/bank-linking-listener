package http

import (
	"demo/bank-linking-listener/internal/delivery/http/http_dto"
	"demo/bank-linking-listener/internal/service"
	"demo/bank-linking-listener/pkg/errorx"
	"demo/bank-linking-listener/pkg/utils"
	"net/http"
	"strconv"

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
	userID := c.MustGet("userID").(uint)
	banks, err := h.bankService.GetBankListByUserID(c.Request.Context(), userID)
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

func (h *BankHandler) GetBankListByUserID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil || userID <= 0 {
		c.JSON(http.StatusBadRequest, utils.ResponseWithMessage(
			utils.ResponseStatusFail, "invalid user id"))
	}

	banks, err := h.bankService.GetBankListByUserID(c.Request.Context(), uint(userID))
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

func (h *BankHandler) CreateBank(c *gin.Context) {
	var req http_dto.BankCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseWithMessage(
			utils.ResponseStatusFail, err.Error()))
		return
	}

	bank, err := h.bankService.CreateBank(c.Request.Context(), *req.ToEntity())
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

// func (h *BankHandler) LinkBank(c *gin.Context) {
// 	userID, err := strconv.Atoi(c.Param("userID"))
// 	if err != nil || userID <= 0 {
// 		c.JSON(http.StatusBadRequest, utils.ResponseWithMessage(
// 			utils.ResponseStatusFail, "invalid user id"))
// 	}
// 	bankCode := c.Param("bankCode")

// 	if err := h.bankService.LinkBank(c.Request.Context(), uint(userID), bankCode); err != nil {
// 		c.JSON(errorx.GetHTTPCode(err), utils.ResponseWithMessage(
// 			utils.ResponseStatusFail, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, utils.ResponseWithMessage(
// 		utils.ResponseStatusSuccess, "bank has been linked"))
// }
