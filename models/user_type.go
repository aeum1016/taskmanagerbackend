package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id" db:"id"`
	Username string    `json:"username" db:"username" binding:"required"`
	Password string    `json:"password" db:"password" binding:"required"`
}