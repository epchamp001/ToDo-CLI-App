package tasks

import (
	"CLITodoApp/internal/entity"
	"fmt"
	"time"
)

func CreateSubtask(description string, parentID int, tasks *[]*entity.Task) error {
	if len(*tasks) == 0 {
		return fmt.Errorf("tasks list is empty")
	}

	var parent *entity.Task
	for _, task := range *tasks {
		if task.ID == parentID {
			parent = task
			break
		}
	}

	if parent == nil {
		return fmt.Errorf("parent task is nil")
	}

	subtaskID := entity.ID.GenerateTaskID()
	subtask := entity.NewTask(subtaskID, time.Now(), description, parent)

	if err := parent.AddSubtask(subtask); err != nil {
		return fmt.Errorf("failed to add subtask: %w", err)
	}

	SortTasksByDeadline(&parent.Subtask)
	return nil
}
