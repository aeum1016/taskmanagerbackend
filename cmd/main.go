package main

import (
	"github.com/aeum1016/taskmanagerbackend/models"
	"github.com/aeum1016/taskmanagerbackend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	connection := models.DBConnection()
	defer connection.Close()

	r := gin.Default()
	
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	r.Use(cors.New(config))
	
	routes.InitRoutes(r)
	r.Run(":8080")
}