package middleware

import (
	"fmt"
	"net/http"

	"github.com/aeum1016/taskmanagerbackend/controllers/session_controller"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println(ctx.Request.Cookies())
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
		fmt.Println(token)
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
