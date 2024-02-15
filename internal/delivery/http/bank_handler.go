package http

import (
	"demo/bank-linking-listener/internal/delivery/http/http_dto"
	"demo/bank-linking-listener/internal/service"
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
	banks, rerr := h.bankService.GetBankListByUserID(c.Request.Context(), userID)
	if rerr != nil {
		c.JSON(utils.GetStatusCode(rerr), utils.ResponseWithMessage(
			utils.ResponseStatusFail, rerr.Message()))
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

	banks, rerr := h.bankService.GetBankListByUserID(c.Request.Context(), uint(userID))
	if rerr != nil {
		c.JSON(utils.GetStatusCode(rerr), utils.ResponseWithMessage(
			utils.ResponseStatusFail, rerr.Message()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseWithData(
		utils.ResponseStatusSuccess, map[string]interface{}{
			"user_id": userID,
			"banks": banks,
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

// func (h *BankHandler) LinkBank(c *gin.Context) {
// 	userID, err := strconv.Atoi(c.Param("userID"))
// 	if err != nil || userID <= 0{
// 		c.JSON(http.StatusBadRequest, utils.ResponseWithMessage(
// 			utils.ResponseStatusFail, "invalid user id"))
// 	}
// 	bankCode := c.Param("bankCode")

// 	if rerr := h.bankService.LinkBank(c.Request.Context(), uint(userID), bankCode); rerr != nil {
// 		c.JSON(utils.GetStatusCode(rerr), utils.ResponseWithMessage(
// 			utils.ResponseStatusFail, rerr.Message()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, utils.ResponseWithMessage(
// 		utils.ResponseStatusSuccess, "bank has been linked"))
// }
