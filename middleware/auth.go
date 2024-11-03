package middleware

import (
	"net/http"
	"os"

	"github.com/aeum1016/taskmanagerbackend/controllers/session_controller"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func AuthMiddleware() gin.HandlerFunc {
	secret := os.Getenv("JWT_PHRASE")

	return func(ctx *gin.Context) {
		token, err := ctx.Request.Cookie("jwt")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Not logged in",
			})
			return
		}

		parsedToken, err := jwt.ParseWithClaims(token.Value, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Failed to parse authentication token",
			})
		} else if claims, ok := parsedToken.Claims.(*jwt.RegisteredClaims); ok {
			parsedUuid, err := uuid.Parse(claims.ID)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "Malformed claim id",
				})		
				return
			}
			foundSession, err := session_controller.FindSessionByID(parsedUuid)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "Not logged in",
				})		
				return
			}
			ctx.Set("uid", foundSession.UID)
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unknown claims type",
			})		
		}
	}
}