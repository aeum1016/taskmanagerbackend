package controllers

import (
	"time"

	"github.com/aeum1016/taskmanagerbackend/models"
	"github.com/gin-gonic/gin"
)

type TaskController interface {
	GetAllTasks() []models.Task
	GetTaskByID(ctx *gin.Context) models.Task
}

func GetAllTasks() []models.Task {
	tasks := []models.Task{
	{
		ID: "123456",
		UID: "12345566",
		Title: "My Task",
		Description: "This is a description for a task 1",
		DueDate: time.Now().AddDate(0, 1, 0),
		Priority: 2,
		EstimateHours: 2,
		Completed: false,
	}, 
	{
		ID: "123457",
		UID: "12345566",
		Title: "My Task 2",
		Description: "This is a description for a task 2",
		DueDate: time.Now().AddDate(0, 1, 1),
		Priority: 3,
		EstimateHours: 4,
		Completed: false,
	},
}

	return tasks
}