package tasks

import (
	"CLITodoApp/internal/entity"
	"fmt"
)

func UpdateTaskStatus(taskID int, status entity.Status, tasks *[]*entity.Task) error {
	for _, task := range *tasks {
		if task.ID == taskID {
			task.SetStatus(status)
			return nil
		}
	}

	return fmt.Errorf("task with ID %d not found", taskID)
}
