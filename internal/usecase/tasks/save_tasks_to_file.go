package tasks

import (
	"CLITodoApp/internal/entity"
	"encoding/json"
	"fmt"
	"os"
)

func SaveTasksToFile(fileName string, tasks *[]*entity.Task) error {
	data, err := json.MarshalIndent(*tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling tasks: %w", err)
	}

	if err := os.WriteFile(fileName, data, 0644); err != nil {
		return fmt.Errorf("error writing tasks to file: %w", err)
	}

	return nil
}
