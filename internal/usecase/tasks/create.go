package tasks

import (
	"CLITodoApp/internal/entity"
	"time"
)

func CreateTask(description string, deadline time.Time, tasks *[]*entity.Task) error {
	taskID := entity.ID.GenerateTaskID()
	task := entity.NewTask(taskID, deadline, description, nil)
	*tasks = append(*tasks, task)
	SortTasksByDeadline(tasks)
	return nil
}
