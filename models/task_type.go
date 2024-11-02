package models

import "time"

type Task struct {
	ID        string    `json:"id" db:"ID"`
	UID       string    `json:"userid" db:"userID"`
	Title     string    `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	DueDate   time.Time `json:"duedate" db:"dueDate"`
	Priority  uint8     `json:"priority" db:"priority"`
	EstimateHours uint8 `json:"estimatehours" db:"hoursEstimate"`
	Tags	[]string			`json:"tags" db:"tags"`
	Completed bool      `json:"completed" db:"completed"`
}

type RecurringTask struct {
	TID string `json:"taskid"`
	Interval time.Time `json:"interval"`
}