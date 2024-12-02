package cmd

import (
	tasks2 "CLITodoApp/internal/usecase/tasks"
	"fmt"
	"github.com/spf13/cobra"
)

var tasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "Show current tasks",
	Long:  "Display a list of current tasks loaded from tasks.json.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := tasks2.LoadTasksFromFile("internal/repository/tasks.json")
		if err != nil {
			fmt.Println("Error loading tasks from file:", err)
			return
		}

		if len(tasks) == 0 {
			fmt.Println("The task list is empty")
			return
		}

		tasks2.ShowAllTasks(tasks)
	},
}

func init() {
	rootCmd.AddCommand(tasksCmd)
}
