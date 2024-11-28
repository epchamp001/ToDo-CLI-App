package tasks

import (
	"CLITodoApp/internal/entity"
	"fmt"
)

func DeleteTask(taskID int, tasks *[]*entity.Task) error {
	if len(*tasks) == 0 {
		return fmt.Errorf("task list is empty")
	}

	for i, task := range *tasks {
		if task.ID == taskID {
			*tasks = append((*tasks)[:i], (*tasks)[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", taskID)
}
