package route

import (
	"demo/bank-linking-listener/internal/delivery/http"
	"demo/bank-linking-listener/internal/service/entity"

	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	Router         *gin.RouterGroup
	UserHandler    *http.UserHandler
	BankHandler    *http.BankHandler
	JWTMiddleware  gin.HandlerFunc
	RoleMiddleware func(...entity.Role) gin.HandlerFunc
}

func (c *RouterConfig) Setup() {
	c.setupGuestRoute()
	c.setupAuthRoute()
}

func (c *RouterConfig) setupGuestRoute() {
	// check health
	{
		c.Router.GET("users/health", c.UserHandler.CheckHealth)
		c.Router.GET("banks/health", c.UserHandler.CheckHealth)
	}

	// user route
	{
		c.Router.POST("users/signup", c.UserHandler.CreateUserAccount)
		c.Router.POST("users/signin", c.UserHandler.SignIn)
	}
}

func (c *RouterConfig) setupAuthRoute() {
	authRoute := c.Router.Group("", c.JWTMiddleware)

	// accept list of role
	adminRoute := authRoute.Group("", c.RoleMiddleware(entity.AdminRole))
	userRoute := authRoute.Group("", c.RoleMiddleware(entity.UserRole))
	customerAdminRoute := authRoute.Group("", c.RoleMiddleware(entity.CustomerRole, entity.AdminRole))

	// user route
	{
		adminRoute.POST("admin/customer-account", c.UserHandler.CreateCustomerAccount) // admin
	}

	// bank route
	{
		userRoute.GET("users/me/banks", c.BankHandler.GetBankListCurrentUser)            // user
		customerAdminRoute.GET("users/:userID/banks", c.BankHandler.GetBankListByUserID) // customer admin
		adminRoute.POST("admin/banks", c.BankHandler.CreateBank)                         // admin
	}
}
