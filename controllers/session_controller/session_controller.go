package session_controller

import (
	"context"
	"time"

	"github.com/aeum1016/taskmanagerbackend/models"
	"github.com/jackc/pgx/v5"
)

type SessionController interface {
	FindSessionByToken(token string) (models.Session, error)
	RemoveExpiredSessions() error
}

func FindSessionByToken(token string) (models.Session, error) {
	db := models.Connection

	query := `SELECT * FROM auth.sessions WHERE "sessionToken" = @token`
	args := pgx.NamedArgs{
		"token": token,
	}

	rows, err := db.Query(context.Background(), query, args)
	if err != nil {
		return models.Session{}, err
	}

	foundSession, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Session])
	if err != nil {
		return models.Session{}, err
	}

	if foundSession.ExpiresAt.Compare(time.Now()) == -1 {
		query := `DELETE FROM auth.sessions WHERE id = @id`
		args := pgx.NamedArgs{
			"id": foundSession.ID,
		}
		_, err := db.Exec(context.Background(), query, args)
		if err != nil {
			return models.Session{}, err
		}
		return models.Session{}, pgx.ErrNoRows
	}

	return foundSession, nil
}

func RemoveExpiredSessions() error {
	db := models.Connection

	_, err := db.Exec(context.Background(), "DELETE FROM auth.sessions WHERE expires < CURRENT_TIMESTAMP")
	if err != nil {
		return err
	}

	return nil
}
