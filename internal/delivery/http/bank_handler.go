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
