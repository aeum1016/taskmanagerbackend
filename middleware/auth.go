package middleware

import (
	"net/http"

	"github.com/aeum1016/taskmanagerbackend/controllers/session_controller"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("authjs.session-token")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Not logged in",
			})		
			return
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