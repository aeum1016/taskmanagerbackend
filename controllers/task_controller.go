package controllers

import (
	"time"

	"github.com/aeum1016/taskmanagerbackend/models"
	"github.com/gin-gonic/gin"
)

type TaskController interface {
	GetAllTasks() []models.Task
	GetTaskByID(ctx *gin.Context) models.Task
	AddTask(ctx *gin.Context) (models.Task, error)
}

var ExampleTasks = []models.Task{
	{
		ID:            "123456",
		UID:           "12345566",
		Title:         "My Task",
		Description:   "This is a description for a task 1",
		DueDate:       time.Now().AddDate(0, 1, 0),
		Priority:      2,
		EstimateHours: 2,
		Completed:     false,
	},
	{
		ID: "123457",
		UID: "12345566",
		Title: "My Task 2",
		Description: "This is a description for a task 2. This description turns out to be reallllllllllllllllllllllllllllllllllly long. And I mean realllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllly long.",
		DueDate: time.Now().AddDate(0, 1, 1),
		Priority: 3,
		EstimateHours: 4,
		Completed:     false,
	},
}

func GetAllTasks() []models.Task {
	return ExampleTasks
}

func AddTask(ctx *gin.Context) (models.Task, error) {
	var newTask models.Task

	if err := ctx.ShouldBind(&newTask); err != nil {
		return models.Task{}, err
	}

	ExampleTasks = append(ExampleTasks, newTask)
	return newTask, nil
}
