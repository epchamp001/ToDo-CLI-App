package entity

import (
	"fmt"
	"time"
)

type HistoryEntry struct {
	ID        int       `json:"id"`
	Task      *Task     `json:"task"`
	Timestamp time.Time `json:"timestamp"`
}

func New(id int, t *Task) (*HistoryEntry, error) {
	if t == nil {
		return nil, fmt.Errorf("cannot create history entry: task is nil")
	}

	if t.Status == Active {
		return nil, fmt.Errorf("cannot create history entry: task has active status")
	}

	return &HistoryEntry{
		ID:        id,
		Task:      t,
		Timestamp: time.Now(),
	}, nil
}

func (h *HistoryEntry) GetID() int {
	return h.ID
}

func (h *HistoryEntry) GetTask() *Task {
	return h.Task
}

func (h *HistoryEntry) GetTimestamp() time.Time {
	return h.Timestamp
}

func (h *HistoryEntry) IsOrderThanAWeek() bool {
	return time.Since(h.Timestamp) > 7*24*time.Hour
}
