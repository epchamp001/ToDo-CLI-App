package tasks

import (
	"CLITodoApp/internal/entity"
	"fmt"
)

func DeleteSubtaskFromTask(parentID int, subtaskID int, tasks *[]*entity.Task) error {
	if len(*tasks) == 0 {
		return fmt.Errorf("task list is empty")
	}

	var parentTask *entity.Task
	for _, task := range *tasks {
		if task.ID == parentID {
			parentTask = task
			break
		}
	}

	if parentTask == nil {
		return fmt.Errorf("parent task with ID %d not found", parentID)
	}

	if err := parentTask.DeleteSubtask(subtaskID); err != nil {
		return fmt.Errorf("failed to delete subtask with %d: %w", subtaskID, err)
	}

	return nil
}
