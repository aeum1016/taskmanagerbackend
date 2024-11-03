package session_controller

import (
	"context"
	"time"

	"github.com/aeum1016/taskmanagerbackend/models"
	"github.com/aeum1016/taskmanagerbackend/util"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type SessionController interface {
	FindSessionByID(ID uuid.UUID) (models.Session, error)
	GetJWTByID(ID uuid.UUID) (string, error)
	GetJWTByUID(UID uuid.UUID) (string, error)
	CreateSession(UID uuid.UUID) (string, error)
	RemoveExpiredSessions() (error)
}

func FindSessionByID(ID uuid.UUID) (models.Session, error) {
	db := models.Connection

	query := `SELECT * FROM public.sessions WHERE id=@id`
  args := pgx.NamedArgs{
    "id": ID,
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
		query := `DELETE FROM public.sessions WHERE id = @id`
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

func GetJWTByID(ID uuid.UUID) (string, error) {
	session, err := FindSessionByID(ID)
	if err != nil {
		return "", err
	}

	return session.Token, nil
}

func GetJWTByUID(UID uuid.UUID) (string, error) {
	db := models.Connection

	query := `SELECT * FROM public.sessions WHERE uid=@uid`
  args := pgx.NamedArgs{
    "uid": UID,
  }
  
	rows, err := db.Query(context.Background(), query, args)
	if err != nil {
		return "", err
	}
	
	foundSession, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Session])
	if err != nil {
		return "", err
	}

	if foundSession.ExpiresAt.Compare(time.Now()) == -1 {
		query := `DELETE FROM public.sessions WHERE id = @id`
		args := pgx.NamedArgs{
			"id": foundSession.ID,
		}
		_, err := db.Exec(context.Background(), query, args)
		if err != nil {
			return "", err
		}
		return "", pgx.ErrNoRows
	}

	return foundSession.Token, nil
}

func CreateSession(UID uuid.UUID) (string, error) {
	db := models.Connection

	// Create session
	id := uuid.New()
	expiryTime := time.Now().Add(time.Hour * 12)

	signedToken, err := util.CreateToken(id.String(), expiryTime)
	if err != nil {
		return "", err
	}

	query := `INSERT INTO public.sessions (id, uid, expires_at, token) VALUES (@id, @uid, @expiresAt, @token)`
  args := pgx.NamedArgs{
    "id": id,
    "uid": UID,
		"expiresAt": expiryTime,
		"token": signedToken,
  }
  
	_, err = db.Exec(context.Background(), query, args)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func RemoveExpiredSessions() (error) {
	db := models.Connection

	_, err := db.Exec(context.Background(), "DELETE FROM public.sessions WHERE expires_at < CURRENT_TIMESTAMP")
	if err != nil {
		return err
	}

	return nil
}