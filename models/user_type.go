package models

import "database/sql"

type UserAuth struct {
	Id                int            `json:"id" db:"id"`
	UserId            int            `json:"uid" db:"userId"`
	Type              string         `json:"type" db:"type"`
	Provider          string         `json:"provider" db:"provider"`
	ProviderAccountId string         `json:"provideraccountid" db:"providerAccountId"`
	RefreshToken      sql.NullString `json:"refreshtoken" db:"refresh_token"`
	AccessToken       sql.NullString `json:"auth" db:"access_token" binding:"required"`
	ExpiresAt         int            `json:"expiresat" db:"expires_at"`
	IdToken           sql.NullString `json:"idtoken" db:"id_token"`
	Scope             sql.NullString `json:"scope" db:"scope"`
	SessionState      sql.NullString `json:"sessionstate" db:"session_state"`
	TokenType         sql.NullString `json:"tokentype" db:"token_type"`
}
