package task_controller

import (
	"context"
	"errors"

	"github.com/aeum1016/taskmanagerbackend/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type TaskController interface {
	GetTasks(ctx *gin.Context) ([]models.Task, error)
	AddTask(ctx *gin.Context) (models.Task, error)
}

type GetTasksByUIDPayload struct {
	UID uuid.UUID `json:"uid"`
}

func GetTasks(ctx *gin.Context) ([]models.Task, error) {
	db := models.Connection

	uid, ok := ctx.Get("uid"); if !ok {
		return []models.Task{}, errors.New("not authenticated") 
	}

	query := `SELECT * FROM public.tasks WHERE uid=@uid`
	args := pgx.NamedArgs{
		"uid": uid,
	}
  rows, err := db.Query(context.Background(), query, args)

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

	uid, ok := ctx.Get("uid"); if !ok {
		return models.Task{}, errors.New("not authenticated") 
	}

	var newTask models.Task

	if err := ctx.ShouldBind(&newTask); err != nil {
		return models.Task{}, err
	}

  query := `INSERT INTO public.tasks (id, uid, title, priority, due_date, description, hours_estimate, tags, completed) VALUES (@id, @uid, @title, @priority, @duedate, @description, @hours, @tags, @completed)`
  args := pgx.NamedArgs{
    "id": uuid.New(),
    "uid": uid,
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
