package cmd

import (
	tasks2 "CLITodoApp/internal/usecase/tasks"
	"fmt"
	"github.com/spf13/cobra"
)

var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "Show tasks with today's deadline",
	Long:  "Show all tasks that have a deadline set to today.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := tasks2.LoadTasksFromFile("internal/repository/tasks.json")
		if err != nil {
			fmt.Println("Error loading tasks from file:", err)
			return
		}

		todayTasks := tasks2.GetTasksWithTodayDeadline(&tasks)

		if len(todayTasks) == 0 {
			fmt.Println("No tasks with today's deadline.")
			return
		}

		tasks2.ShowAllTasks(todayTasks)
	},
}

func init() {
	rootCmd.AddCommand(todayCmd)
}
