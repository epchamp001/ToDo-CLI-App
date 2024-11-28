package entity

import (
	"fmt"
	"time"
)

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
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	Parent      *Task     `json:"-"`
	Subtask     []*Task   `json:"subtask,omitempty"`
}

func NewTask(id int, deadline time.Time, description string, parent *Task) *Task {
	createdAt := time.Now()
	dl := deadline
	if parent != nil {
		createdAt = parent.CreatedAt
		dl = parent.Deadline
	}

	return &Task{
		ID:          id,
		CreatedAt:   createdAt,
		Deadline:    dl,
		Subtask:     nil,
		Description: description,
		Status:      Active,
		Parent:      parent,
	}
}

func (t *Task) GetStatus() Status {
	return t.Status
}

func (t *Task) SetStatus(status Status) {
	if status == Done || status == Unnecessary {
		for _, subtask := range t.Subtask {
			subtask.SetStatus(status)
		}
	}

	t.Status = status
}

func (t *Task) GetDeadline() time.Time {
	return t.Deadline
}

func (t *Task) UpdateDeadline(deadline time.Time) {
	t.Deadline = deadline
}

func (t *Task) IsDeadlineExpired() bool {
	return time.Now().After(t.Deadline)
}

func (t *Task) AddSubtask(subtask *Task) error {
	if subtask == nil {
		return fmt.Errorf("subtask is nil")
	}

	if t.Subtask == nil {
		t.Subtask = []*Task{}
	}

	for _, st := range t.Subtask {
		if st.Description == subtask.Description && st.Deadline == subtask.Deadline {
			return fmt.Errorf("subtask already exists")
		}
	}

	t.Subtask = append(t.Subtask, subtask)
	subtask.Parent = t

	return nil
}

func (t *Task) DeleteSubtask(id int) error {
	if t.Subtask == nil {
		return fmt.Errorf("subtask list is empty")
	}
	for i, st := range t.Subtask {
		if st.ID == id {
			st.Parent = nil
			t.Subtask = append(t.Subtask[:i], t.Subtask[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("subtask with ID %d not found", id)
}

func (t *Task) GetSubtasks() []*Task {
	if len(t.Subtask) == 0 {
		t.Subtask = []*Task{}
	}
	return t.Subtask
}

func (t *Task) UpdateDescription(description string) {
	t.Description = description
}

func (t *Task) GetDescription() string {
	return t.Description
}

func (t *Task) GetParent() *Task {
	return t.Parent
}

func (t *Task) GetCreatedAt() time.Time {
	return t.CreatedAt
}

func (t *Task) HaveSubtask() bool {
	return len(t.Subtask) > 0
}

func (t *Task) GetSubtask(id int) (*Task, error) {
	for _, st := range t.Subtask {
		if st.ID == id {
			return st, nil
		}
	}
	return nil, fmt.Errorf("subtask with ID %d not found", id)
}
