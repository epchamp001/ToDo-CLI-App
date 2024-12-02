package cmd

import (
	"CLITodoApp/internal/entity"
	tasks2 "CLITodoApp/internal/usecase/tasks"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	taskID int
	status string
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Update the status of a task",
	Long:  "Update the status of a task by specifying its ID and the new status.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := tasks2.LoadTasksFromFile("internal/repository/tasks.json")
		if err != nil {
			fmt.Println("Error loading tasks from file:", err)
			return
		}

		newStatus := entity.Status(status)

		err = tasks2.UpdateTaskStatus(taskID, newStatus, &tasks)
		if err != nil {
			fmt.Println("Error updating task status:", err)
			return
		}

		if err := tasks2.SaveTasksToFile("internal/repository/tasks.json", &tasks); err != nil {
			fmt.Println("Error saving tasks:", err)
			return
		}

		tasks2.ShowAllTasks(tasks)
	},
}

func init() {
	statusCmd.Flags().IntVarP(&taskID, "id", "i", 0, "ID of the task to update")
	statusCmd.Flags().StringVarP(&status, "status", "s", "", "New status for the task (active, done, unnecessary)")

	statusCmd.MarkFlagRequired("id")
	statusCmd.MarkFlagRequired("status")

	rootCmd.AddCommand(statusCmd)
}
