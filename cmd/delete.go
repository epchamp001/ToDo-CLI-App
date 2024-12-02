package cmd

import (
	del "CLITodoApp/cmd/delete"
	tasks2 "CLITodoApp/internal/usecase/tasks"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	parentID int
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Long:  "Delete a task by specifying its ID.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := tasks2.LoadTasksFromFile("internal/repository/tasks.json")
		if err != nil {
			fmt.Println("Error loading tasks from file:", err)
			return
		}

		err = tasks2.DeleteTask(parentID, &tasks)
		if err != nil {
			fmt.Println("Error deleting task:", err)
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
	deleteCmd.Flags().IntVarP(&parentID, "parentid", "p", 0, "ID of the task to delete")
	deleteCmd.MarkFlagRequired("parentid")

	rootCmd.AddCommand(deleteCmd)
	deleteCmd.AddCommand(del.SubtaskDeleteCmd)
}
