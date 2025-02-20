package middlewares

import (
	"mailmind-api/pkg/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		userRole, exists := ctx.Get("role")
		if !exists {
			_ = ctx.Error(utils.NewCustomError(http.StatusUnauthorized, "Unauthorized: No role found"))
			ctx.Abort()
			return
		}

		roleStr, ok := userRole.(string)
		if !ok {
			_ = ctx.Error(utils.NewCustomError(http.StatusInternalServerError, "Invalid role format"))
			ctx.Abort()
			return
		}

		for _, role := range allowedRoles {
			if strings.EqualFold(role, roleStr) {

				ctx.Next()
				return
			}
		}

		_ = ctx.Error(utils.NewCustomError(http.StatusForbidden, "Forbidden: Access denied"))
		ctx.Abort()
	}
}
