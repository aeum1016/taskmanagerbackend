package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/aeum1016/taskmanagerbackend/controllers/session_controller"
	"github.com/gin-gonic/gin"
)

func AuthUserMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("__Secure-authjs.session-token")
		if err != nil {
			token, err = ctx.Cookie("authjs.session-token")
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "Not logged in",
				})
				return
			}
		}
		session, err := session_controller.FindSessionByToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Not logged in",
			})
			return
		}
		ctx.Set("uid", session.UID)
		ctx.Next()
	}
}

func AuthAdminMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Values("Authentication")
		if len(token) == 0 || token[0] != fmt.Sprintf("Basic %s", os.Getenv("AUTH_ADMIN_PASSWORD")) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Not admin user",
			})
			return
		}
		ctx.Next()
	}
}
