package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aeum1016/taskmanagerbackend/controllers/session_controller"
	"github.com/aeum1016/taskmanagerbackend/controllers/task_controller"
	"github.com/aeum1016/taskmanagerbackend/models"
	"github.com/aeum1016/taskmanagerbackend/routes"
	"github.com/aeum1016/taskmanagerbackend/util"
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

	util.Schedule(context.Background(), 24 * time.Hour, 0, func(v time.Time){
		task_controller.RemoveCompletedTasks()
	})
	util.Schedule(context.Background(), 30 * time.Minute, 0, func(v time.Time){
		session_controller.RemoveExpiredSessions()
	})

	r := gin.Default()
	
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowCredentials = true
	r.Use(cors.New(config))
	
	routes.InitRoutes(r)
	r.Run(":8080")
}