package session_controller

import (
	"context"
	"os"
	"time"

	"github.com/aeum1016/taskmanagerbackend/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type SessionController interface {
	CreateSession(UID uuid.UUID) (string, error)
}

func CreateSession(UID uuid.UUID) (string, error) {
	db := models.Connection

	// Create session
	id := uuid.New()
	expiryTime := time.Now().Add(time.Hour * 12)
	query := `INSERT INTO public.sessions (id, uid, expires_at) VALUES (@id, @uid, @expiresAt)`
  args := pgx.NamedArgs{
    "id": id,
    "uid": UID,
		"expiresAt": expiryTime,
  }
  
	_, err := db.Exec(context.Background(), query, args)

	// Create the Claims
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expiryTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID: id.String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	hash := os.Getenv("JWT_PHRASE")
	signingKey := []byte(hash)

	ss, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return ss, nil
}