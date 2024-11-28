package tasks

import (
	"CLITodoApp/internal/entity"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func LoadTasksFromFile(filename string) ([]*entity.Task, error) {
	tasks := make([]*entity.Task, 0)

	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return tasks, nil // Файл не существует, возвращаем пустой список
		}
		return nil, fmt.Errorf("error reading tasks file: %w", err)
	}

	if len(data) == 0 {
		return tasks, nil // Пустой файл, возвращаем пустой список
	}

	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("error unmarshalling tasks: %w", err)
	}

	for _, task := range tasks {
		if task.IsDeadlineExpired() {
			task.UpdateDeadline(time.Now())
		}
	}
	return tasks, nil
}
