package task

import (
	"net/http"

	"github.com/aeum1016/taskmanagerbackend/controllers/task_controller"
	"github.com/aeum1016/taskmanagerbackend/middleware"
	"github.com/gin-gonic/gin"
)

func InitTaskRoutes(r *gin.Engine) {
	tr := r.Group("/task")

	tr.Use(middleware.AuthMiddleware()) 
	{
		tr.GET("", getTasks())
		tr.POST("", addTask())
	}
}

func getTasks() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tasks, err := task_controller.GetTasks(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, tasks)
	}
}

func addTask() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		task, err := task_controller.AddTask(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, task)
	}
}
