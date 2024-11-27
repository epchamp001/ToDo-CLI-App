package entity

import (
	"time"
)

type HistoryEntry struct {
	ID        int       `json:"id"`
	Task      *Task     `json:"task"`
	Timestamp time.Time `json:"timestamp"`
}
