package user

import (
	"net/http"

	"github.com/aeum1016/taskmanagerbackend/controllers/user_controller"
	"github.com/aeum1016/taskmanagerbackend/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(r *gin.Engine) {
	tr := r.Group("/user")

	tr.Use(middleware.AuthMiddleware())
	{
		tr.GET("/token", getUserAuth())
	}
}

func getUserAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth, err := user_controller.GetUserAuth(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, auth)
	}
}
