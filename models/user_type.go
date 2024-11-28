package models

type UserAuth struct {
	Auth string `json:"auth" db:"access_token" binding:"required"`
}
