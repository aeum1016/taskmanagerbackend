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
  UpdateTask(ctx *gin.Context) (models.Task, error)
  RemoveCompletedTasks(ctx *gin.Context) error
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

func UpdateTask(ctx *gin.Context) (models.Task, error) {
  db := models.Connection

  _, ok := ctx.Get("uid"); if !ok {
		return models.Task{}, errors.New("not authenticated") 
	}

	var updatedTask models.Task
	if err := ctx.ShouldBind(&updatedTask); err != nil {
		return models.Task{}, err
	}

  query := `UPDATE public.tasks 
            SET (title, priority, due_date, description, hours_estimate, tags, completed) 
            = (@title, @priority, @duedate, @description, @hours, @tags, @completed)
            WHERE "id" = @id`
  args := pgx.NamedArgs{
    "id": updatedTask.ID,
    "title": updatedTask.Title,
    "priority": updatedTask.Priority,
    "duedate": updatedTask.DueDate,
    "description": updatedTask.Description,
    "hours": updatedTask.EstimateHours,
    "tags": updatedTask.Tags,
    "completed": updatedTask.Completed,
  }

  _, err := db.Exec(context.Background(), query, args)
  if err != nil {
    return models.Task{}, err 
  }

	return updatedTask, nil
}

func RemoveCompletedTasks() error {
  db := models.Connection

	_, err := db.Exec(context.Background(), "DELETE FROM public.tasks WHERE completed = true")
	if err != nil {
		return err
	}

	return nil
}
