package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID         uuid.UUID    `json:"id" db:"id"`
	UID        uuid.UUID    `json:"userid" db:"uid"`
	ExpiresAt time.Time `json:"expiresat" db:"expires_at"`
	Token string `json:"token" db:"token"`
}