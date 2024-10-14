package models

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password" binding:"required"`
}