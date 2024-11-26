package entity

import "time"

type Status string

const (
	Active      Status = "active"
	Done        Status = "done"
	Unnecessary Status = "unnecessary"
)

type Task struct {
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Deadline    time.Time `json:"deadline"`
	Subtask     []Subtask `json:"subtask"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
}
