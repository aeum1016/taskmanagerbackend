package task

import (
	"net/http"

	"github.com/aeum1016/taskmanagerbackend/controllers"
	"github.com/gin-gonic/gin"
)

func InitTaskRoutes(r *gin.Engine) {
	tr := r.Group("/task")

	tr.GET("/", getAllTasks())
	tr.POST("/", addTask())

}

func getAllTasks() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, controllers.GetAllTasks())
	}
}

func addTask() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		task, err := controllers.AddTask(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
		}
		if err == nil {
			ctx.JSON(http.StatusOK, task)
		}
	}
}
