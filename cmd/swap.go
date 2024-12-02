package cmd

import (
	tasks2 "CLITodoApp/internal/usecase/tasks"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	taskID1 int
	taskID2 int
)

var swapCmd = &cobra.Command{
	Use:   "swap",
	Short: "Swap two tasks",
	Long:  "Swap the positions of two tasks in the list by specifying their IDs.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := tasks2.LoadTasksFromFile("internal/repository/tasks.json")
		if err != nil {
			fmt.Println("Error loading tasks from file:", err)
			return
		}

		err = tasks2.ReplaceTasks(taskID1, taskID2, &tasks)
		if err != nil {
			fmt.Println("Error swapping tasks:", err)
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
	swapCmd.Flags().IntVarP(&taskID1, "id1", "a", 0, "ID of the first task to swap")
	swapCmd.Flags().IntVarP(&taskID2, "id2", "b", 0, "ID of the second task to swap")

	swapCmd.MarkFlagRequired("id1")
	swapCmd.MarkFlagRequired("id2")

	rootCmd.AddCommand(swapCmd)
}
