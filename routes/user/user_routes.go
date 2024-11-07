package user

import (
	"net/http"

	"github.com/aeum1016/taskmanagerbackend/controllers/user_controller"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(r *gin.Engine) {
	tr := r.Group("/user")

	tr.POST("/login", loginUser())
	tr.POST("/create", createUser())
}

func createUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := user_controller.CreateUser(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{})
	}
}

func loginUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := user_controller.LoginUser(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{})
	}
}