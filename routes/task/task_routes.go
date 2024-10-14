package task

import (
	"net/http"

	"github.com/aeum1016/taskmanagerbackend/controllers"
	"github.com/gin-gonic/gin"
)

func InitTaskRoutes(r *gin.Engine) {
	tr := r.Group("/task") 
	
	tr.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, controllers.GetAllTasks())
	})
}