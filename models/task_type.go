package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID            uuid.UUID   `json:"id" db:"id"`
	UID           json.Number `json:"uid" db:"uid"`
	Title         string      `json:"title" db:"title" binding:"required"`
	Description   string      `json:"description" db:"description"`
	DueDate       time.Time   `json:"duedate" db:"due_date"`
	Priority      uint8       `json:"priority" db:"priority"`
	EstimateHours json.Number `json:"estimatehours" db:"hours_estimate"`
	Tags          []string    `json:"tags" db:"tags"`
	Completed     bool        `json:"completed" db:"completed"`
}

type RecurringTask struct {
	TID      string    `json:"taskid"`
	Interval time.Time `json:"interval"`
}
