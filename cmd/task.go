package cmd

import (
	"CLITodoApp/internal/entity"
	tasks2 "CLITodoApp/internal/usecase/tasks"
	"fmt"
	"github.com/spf13/cobra"
)

var infoTaskID int

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Show detailed task information",
	Long:  "Display detailed information about a specific task by its ID from the current task list.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := tasks2.LoadTasksFromFile("internal/repository/tasks.json")
		if err != nil {
			fmt.Println("Error loading tasks from file:", err)
			return
		}

		var infoTask *entity.Task
		for _, task := range tasks {
			if task.ID == infoTaskID {
				infoTask = task
				break
			}
		}

		if infoTask == nil {
			fmt.Printf("Task with ID %d not found.\n", infoTaskID)
			return
		}

		tasks2.ShowTask(infoTask)
	},
}

func init() {
	taskCmd.Flags().IntVarP(&infoTaskID, "id", "i", 0, "Task ID to display details")
	taskCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(taskCmd)
}
