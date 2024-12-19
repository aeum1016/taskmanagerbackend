package session

import (
	"net/http"

	"github.com/aeum1016/taskmanagerbackend/controllers/session_controller"
	"github.com/aeum1016/taskmanagerbackend/middleware"
	"github.com/gin-gonic/gin"
)

func InitSessionRoutes(r *gin.Engine) {
	srAdmin := r.Group("/session/admin")

	srAdmin.Use(middleware.AuthAdminMiddleware())
	{
		srAdmin.GET("/removeExpired", removeExpiredSessions())
	}
}

func removeExpiredSessions() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := session_controller.RemoveExpiredSessions()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"response": "Successfully removed expired sessions",
		})
	}
}
