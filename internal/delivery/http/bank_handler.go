package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BankHandler struct {
	// Noncompliant
}

func NewBankHandler() *BankHandler {
	return &BankHandler{}
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
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}