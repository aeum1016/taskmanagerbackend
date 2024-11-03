package user_controller

import (
	"context"

	"github.com/aeum1016/taskmanagerbackend/controllers/session_controller"
	"github.com/aeum1016/taskmanagerbackend/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type UserController interface {
	LoginUser(ctx *gin.Context) (string, error)
	CreateUser(ctx *gin.Context) (string, error)
}

type LoginPayload struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginUser(ctx *gin.Context) (string, error) {
	db := models.Connection

	var user LoginPayload
	if err := ctx.ShouldBind(&user); err != nil {
		return "", err
	}

	query := `SELECT * FROM public.users WHERE username = @username AND password = @password` 
	args := pgx.NamedArgs{
		"username": user.Username,
		"password": user.Password,
	}

  rows, err := db.Query(context.Background(), query, args)
	if err != nil {
		return "", err
	}

	foundUser, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])
	if err != nil {
		return "", err
	}

	userID, err := uuid.Parse(foundUser.ID)
	if err != nil {
		return "", err
	}

	return session_controller.CreateSession(userID)
}

func CreateUser(ctx *gin.Context) (string, error) {
	db := models.Connection

	var user LoginPayload
	if err := ctx.ShouldBind(&user); err != nil {
		return "", err
	}

	givenID := uuid.New()
	query := `INSERT INTO public.users(id, username, password) VALUES (@id, @username, @password)`
	args := pgx.NamedArgs{
		"id": givenID,
		"username": user.Username,
		"password": user.Password,
	}

	_, err := db.Exec(context.Background(), query, args)
	if err != nil {
		return "", err
	}

	return session_controller.CreateSession(givenID)
}