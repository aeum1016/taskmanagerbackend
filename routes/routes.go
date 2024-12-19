package routes

import (
	"github.com/aeum1016/taskmanagerbackend/routes/session"
	"github.com/aeum1016/taskmanagerbackend/routes/task"
	"github.com/aeum1016/taskmanagerbackend/routes/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	task.InitTaskRoutes(r)
	user.InitUserRoutes(r)
	session.InitSessionRoutes(r)
}
