package middleware

import (
	"net/http"
	"strings"

	"github.com/ncondezo/final/pkg/security"
	"github.com/ncondezo/final/pkg/web"

	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if !strings.HasPrefix(token, "Bearer ") {
			ctx.AbortWithStatusJSON(
				http.StatusForbidden,
				web.ErrorResponse{
					http.StatusForbidden,
					"Se requiere un token de acceso.",
				})
			return
		}
		token = strings.TrimPrefix(token, "Bearer ")
		_, err := security.ValidateToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusForbidden,
				web.ErrorResponse{
					http.StatusForbidden,
					"Token de acceso inválido.",
				})
			return
		}
		ctx.Next()
		return
	}
}
