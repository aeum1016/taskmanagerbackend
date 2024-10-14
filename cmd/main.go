package main

import (
	"github.com/gin-gonic/gin"

	"github.com/aeum1016/taskmanagerbackend/routes"
)

func main() {
	r := gin.Default()
	routes.InitRoutes(r)
	r.Run(":8080")
}