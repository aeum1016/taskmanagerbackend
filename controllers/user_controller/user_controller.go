package user_controller

import (
	"context"
	"errors"

	"github.com/aeum1016/taskmanagerbackend/models"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type TaskController interface {
	GetUserAuth(ctx *gin.Context) (models.UserAuth, error)
}

func GetUserAuth(ctx *gin.Context) (models.UserAuth, error) {
	db := models.Connection

	uid, ok := ctx.Get("uid")
	if !ok {
		return models.UserAuth{}, errors.New("not authenticated")
	}

	query := `SELECT * FROM public.accounts WHERE userid=@uid`
	args := pgx.NamedArgs{
		"uid": uid,
	}
	rows, err := db.Query(context.Background(), query, args)

	if err != nil {
		return models.UserAuth{}, err
	}

	auth, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.UserAuth])

	if err != nil {
		return models.UserAuth{}, err
	}

	return auth, nil
}
