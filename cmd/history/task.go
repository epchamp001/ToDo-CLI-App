package history

import (
	"CLITodoApp/internal/entity"
	"CLITodoApp/internal/usecase/tasks"
	"fmt"
	"github.com/spf13/cobra"
)

var entryID int

var TaskCmd = &cobra.Command{
	Use:   "task",
	Short: "Show detailed task information from history",
	Long:  "Display detailed information about a specific task in history by its entry ID.",
	Run: func(cmd *cobra.Command, args []string) {
		history := entity.NewHistory(1)

		if err := history.LoadFromFile("internal/repository/history.json"); err != nil {
			fmt.Println("Error loading history:", err)
			return
		}

		entry, err := history.GetEntry(entryID)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		task := entry.GetTask()
		if task == nil {
			fmt.Printf("Task not found for entry ID %d\n", entryID)
			return
		}

		tasks.ShowTask(task)
	},
}

func init() {
	TaskCmd.Flags().IntVarP(&entryID, "id", "i", 0, "Entry ID of the task in history")
	TaskCmd.MarkFlagRequired("id")
}
