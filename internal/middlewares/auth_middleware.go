package middlewares

import (
	"mailmind-api/config"
	"mailmind-api/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(config *config.AppConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		accessToken, err := ctx.Cookie("access_token")
		if err != nil {
			_ = ctx.Error(utils.NewCustomError(http.StatusUnauthorized, "No access token found"))
			ctx.Abort()
			return
		}

		claims, err := utils.ValidateToken(accessToken, config.AccessTokenSecret)
		if err != nil {
			_ = ctx.Error(utils.NewCustomError(http.StatusUnauthorized, "Invalid or expired access token"))
			ctx.Abort()
			return
		}

		ctx.Set("user_id", claims.ID)
		ctx.Next()
	}
}
