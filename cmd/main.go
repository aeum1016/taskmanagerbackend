package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/aeum1016/taskmanagerbackend/routes"
)

func main() {
	r := gin.Default()
	
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	r.Use(cors.New(config))
	
	routes.InitRoutes(r)
	r.Run(":8080")
}