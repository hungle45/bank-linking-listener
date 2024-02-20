package http

import (
	"demo/bank-linking-listener/config"
	"demo/bank-linking-listener/internal/delivery/http/middleware"
	"demo/bank-linking-listener/internal/service"
	"demo/bank-linking-listener/internal/service/entity"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	cfg         *config.Config
	userService service.UserService
	bankService service.BankService
}

func NewController(cfg *config.Config, userService service.UserService, bankService service.BankService) *Controller {
	return &Controller{
		cfg:         cfg,
		userService: userService,
		bankService: bankService,
	}
}

func (controller *Controller) Routes(r *gin.Engine) {
	jwtMiddleware := middleware.JWTMiddleware()
	roleMiddlewareFn := middleware.RoleMiddleware(controller.userService)
	adminRoleMiddleware := roleMiddlewareFn(entity.AdminRole)
	userRoleMiddleware := roleMiddlewareFn(entity.UserRole)
	customerAdminRoleMiddleware := roleMiddlewareFn(entity.CustomerRole, entity.AdminRole)

	v1 := r.Group("/v1")
	{
		v1.POST("users/signup", controller.CreateUserAccount)
		v1.POST("users/signin", controller.SignIn)

		authRoute := v1.Group("", jwtMiddleware)
		{
			adminRoute := authRoute.Group("", adminRoleMiddleware)
			{
				adminRoute.POST("admin/banks", controller.CreateBank)
				adminRoute.POST("admin/customer-account", controller.CreateCustomerAccount)
			}

			userRoute := authRoute.Group("", userRoleMiddleware)
			{
				userRoute.GET("users/me/banks", controller.GetBankListCurrentUser)
			}

			customerAdminRoute := authRoute.Group("", customerAdminRoleMiddleware)
			{
				customerAdminRoute.GET("users/:userID/banks", controller.GetBankListByUserID)
			}
		}
	}
}
