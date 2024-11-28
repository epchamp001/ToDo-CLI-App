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

	for i, subtask := range parentTask.Subtask {
		if subtask.ID == subtaskID {
			parentTask.Subtask = append(parentTask.Subtask[:i], parentTask.Subtask[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("subtask with ID %d not found under parent task with ID %d", subtaskID, parentID)
}
