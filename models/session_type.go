package models

import "time"

type Session struct {
	ID         string    `json:"id" db:"id"`
	UID        string    `json:"userid" db:"userID"`
	ExpiresAt time.Time `json:"expiresat" db:"expiresAt"`
}