package task

import (
	"net/http"

	"github.com/aeum1016/taskmanagerbackend/controllers/task_controller"
	"github.com/gin-gonic/gin"
)

func InitTaskRoutes(r *gin.Engine) {
	tr := r.Group("/task")

	tr.GET("", getAllTasks())
	tr.POST("", addTask())
}

func getAllTasks() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tasks, err := task_controller.GetAllTasks()
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
		}
		ctx.JSON(http.StatusOK, task)
	}
}
