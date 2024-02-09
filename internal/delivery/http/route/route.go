package route

import (
	"demo/bank-linking-listener/internal/delivery/http"

	"github.com/gin-gonic/gin"
)

type routerConfig struct {
	router      *gin.RouterGroup
	userHandler *http.UserHandler
	bankHandler *http.BankHandler
}

func NewRouterConfig(router *gin.RouterGroup, userHandler *http.UserHandler, bankHandler *http.BankHandler) *routerConfig {
	return &routerConfig{
		router:      router,
		userHandler: userHandler,
		bankHandler: bankHandler,
	}
}

func (c *routerConfig) Setup() {
	c.setupUserRoute()
	c.setupBankRoute()
}

func (c *routerConfig) setupUserRoute() {
	userRouteGroup := c.router.Group("/user")

	// guest route
	{
		userRouteGroup.GET("/health", c.userHandler.CheckHealth)
	}
}

func (c *routerConfig) setupBankRoute() {
	bankRouteGroup := c.router.Group("/bank")

	// guest route
	{
		bankRouteGroup.GET("/health", c.bankHandler.CheckHealth)
	}
}
