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
	c.setupGuestRoute()
	c.setupAuthRoute()
}

func (c *routerConfig) setupGuestRoute() {
	// check health
	{
		c.router.GET("users/health", c.userHandler.CheckHealth)
		c.router.GET("banks/health", c.userHandler.CheckHealth)
	}

	// user route
	{
		c.router.POST("users/signup", c.userHandler.CreateUserAccount)
		c.router.POST("users/signin", c.userHandler.SignIn)
	}
}

func (c *routerConfig) setupAuthRoute() {
	// user route
	{
		c.router.POST("admin/customer-account", c.userHandler.CreateCustomerAccount) // admin
	}
	
	// bank route
	{
		c.router.GET("users/me/banks", c.bankHandler.GetBankListCurrentUser) // user
		c.router.GET("users/:userID/banks", c.bankHandler.GetBankListByUserID) // system admin
		// c.router.POST("admin/banks", c.bankHandler.CreateBank) // admin
	}
}
