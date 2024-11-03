package task_controller

import (
	"context"

	"github.com/aeum1016/taskmanagerbackend/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type TaskController interface {
	GetAllTasks() ([]models.Task, error)
	AddTask(ctx *gin.Context) (models.Task, error)
}

func GetAllTasks() ([]models.Task, error) {
	db := models.Connection
  rows, err := db.Query(context.Background(), "SELECT * FROM public.tasks")
  if err != nil {
    return []models.Task{}, err
  }
  
  tasks, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Task])

  if err != nil {
    return []models.Task{}, err
  }

  return tasks, nil
}

func AddTask(ctx *gin.Context) (models.Task, error) {
  db := models.Connection

	var newTask models.Task

	if err := ctx.ShouldBind(&newTask); err != nil {
		return models.Task{}, err
	}


  query := `INSERT INTO public.tasks (id, uid, title, priority, due_date, description, hours_estimate, tags, completed) VALUES (@id, @uid, @title, @priority, @duedate, @description, @hours, @tags, @completed)`
  args := pgx.NamedArgs{
    "id": uuid.New(),
    "uid": newTask.UID,
    "title": newTask.Title,
    "priority": newTask.Priority,
    "duedate": newTask.DueDate,
    "description": newTask.Description,
    "hours": newTask.EstimateHours,
    "tags": newTask.Tags,
    "completed": newTask.Completed,
  }

  _, err := db.Exec(context.Background(), query, args)

  if err != nil {
    return models.Task{}, err 
  }

	return newTask, nil
}
