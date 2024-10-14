package models

import "time"

type Task struct {
	ID        string    `json:"id"`
	UID       string    `json:"userid"`
	Title     string    `json:"title"`
	DueDate   time.Time `json:"duedate"`
	Priority  uint8     `json:"priority"`
	Completed bool      `json:"completed"`
}

type RecurringTask struct {
	TID string `json:"taskid"`
	Interval time.Time `json:"interval"`
}