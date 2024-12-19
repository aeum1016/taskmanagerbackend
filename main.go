package main

import (
	"fmt"

	"github.com/aeum1016/taskmanagerbackend/models"
	"github.com/aeum1016/taskmanagerbackend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file, ", err)
	}

	connection := models.DBConnection()
	defer connection.Close()

	r := gin.Default()

	config := cors.DefaultConfig()

	config.AllowOrigins = []string{"http://localhost", "https://tm-frontend-218016110927.us-central1.run.app"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	routes.InitRoutes(r)
	r.Run(":8080")
}
