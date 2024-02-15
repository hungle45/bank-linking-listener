package middleware

import (
	"demo/bank-linking-listener/internal/service"
	"demo/bank-linking-listener/internal/service/entity"
	"demo/bank-linking-listener/pkg/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				utils.ResponseWithMessage(utils.ResponseStatusFail, "bad header value given"))
		}

		token := strings.Split(header, " ")
		if len(token) != 2 || token[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				utils.ResponseWithMessage(utils.ResponseStatusFail, "invalid formatted authorization header"))
		}

		userEmail, err := utils.ParseToken(token[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				utils.ResponseWithMessage(utils.ResponseStatusFail, "invalid token"))
		}

		c.Set("userEmail", userEmail)
		c.Next()
	}
}

func RoleMiddleware(userService service.UserService) func(...entity.Role) gin.HandlerFunc {
	return func(acceptedRole ...entity.Role) gin.HandlerFunc {
		return func(c *gin.Context) {
			userEmail := c.MustGet("userEmail").(string)
			user, rerr := userService.GetByEmail(c, userEmail)
			if rerr != nil {
				c.AbortWithStatusJSON(utils.GetStatusCode(rerr),
					utils.ResponseWithMessage(utils.ResponseStatusFail, rerr.Message()))
			}
			fmt.Println(acceptedRole, user.Role)
			for _, role := range acceptedRole {
				if role == user.Role {
					c.Next()
					return
				}
			}

			c.AbortWithStatusJSON(http.StatusForbidden,
				utils.ResponseWithMessage(utils.ResponseStatusFail, "forbidden"))
		}
	}
}
