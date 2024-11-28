package tasks

import (
	"CLITodoApp/internal/entity"
	"fmt"
	"time"
)

func CreateSubtask(description string, deadline time.Time, tasks *[]*entity.Task, parent *entity.Task) error {
	if parent == nil {
		return fmt.Errorf("parent task is nil")
	}

	subtaskID := entity.ID.GenerateTaskID()

	subtask := entity.NewTask(subtaskID, deadline, description, parent)

	if parent.Subtask == nil {
		parent.Subtask = []*entity.Task{}
	}
	parent.Subtask = append(parent.Subtask, subtask)
	SortTasksByDeadline(&parent.Subtask)
	return nil
}
