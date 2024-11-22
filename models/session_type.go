package models

import (
	"time"
)

type Session struct {
	ID         	int    				`json:"id" db:"id"`
	UID        	int    				`json:"uid" db:"userId"`
	ExpiresAt 	time.Time 	 	`json:"expires" db:"expires"`
	Token 			string 			 	`json:"token" db:"sessionToken"`
}