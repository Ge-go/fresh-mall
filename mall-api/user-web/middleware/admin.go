package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"

	jwtmodel "mall-api/user-web/global/jwt"
)

func IsAdminAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, _ := ctx.Get("claims")
		customClaims := claims.(*jwtmodel.CustomClaims)

		if customClaims.AuthorityId != 2 {
			ctx.JSON(http.StatusForbidden, gin.H{
				"msg": "no permission",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
