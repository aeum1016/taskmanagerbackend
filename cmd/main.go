package main

import (
	"fmt"
	"time"

	"github.com/aeum1016/taskmanagerbackend/controllers/session_controller"
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

	ticker := time.NewTicker(30 * time.Minute)
	quit := make(chan bool)
	defer func() {
		ticker.Stop()
		quit <- true
	}()

	go func(quit chan bool) {
			for {
				 select {
					case <- ticker.C:
						session_controller.RemoveExpiredSessions()
					case <- quit:
						ticker.Stop()
						return
					}
			}
	 }(quit)

	r := gin.Default()
	
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	r.Use(cors.New(config))
	
	routes.InitRoutes(r)
	r.Run(":8080")
}