package routes

import (
	"github.com/aeum1016/taskmanagerbackend/routes/task"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	task.InitTaskRoutes(r)
}