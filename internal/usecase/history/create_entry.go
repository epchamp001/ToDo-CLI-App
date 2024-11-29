package history

import (
	"CLITodoApp/internal/entity"
	"fmt"
)

func CreateEntryInHistory(task *entity.Task, history *entity.History) error {
	if task == nil {
		return fmt.Errorf("task is nil")
	}
	if history == nil {
		return fmt.Errorf("history is nil")
	}

	entryID := entity.ID.GenerateHistoryEntryID()
	newEntry, err := entity.NewEntry(entryID, task)
	if err != nil {
		return fmt.Errorf("could not create new entry: %w", err)
	}

	if err := history.AddEntry(newEntry); err != nil {
		return fmt.Errorf("could not add new entry to the history: %w", err)
	}
	return nil

}
