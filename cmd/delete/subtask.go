package delete

import (
	tasks2 "CLITodoApp/internal/usecase/tasks"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	parentID  int
	subtaskID int
)

var SubtaskDeleteCmd = &cobra.Command{
	Use:   "subtask",
	Short: "Delete a subtask",
	Long:  "Delete a subtask from a parent task by specifying the parent ID and subtask ID.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := tasks2.LoadTasksFromFile("internal/repository/tasks.json")
		if err != nil {
			fmt.Println("Error loading tasks from file:", err)
			return
		}

		err = tasks2.DeleteSubtaskFromTask(parentID, subtaskID, &tasks)
		if err != nil {
			fmt.Println("Error deleting subtask:", err)
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
	SubtaskDeleteCmd.Flags().IntVarP(&parentID, "parentid", "p", 0, "ID of the parent task")
	SubtaskDeleteCmd.Flags().IntVarP(&subtaskID, "subtaskid", "s", 0, "ID of the subtask to delete")

	SubtaskDeleteCmd.MarkFlagRequired("parentid")
	SubtaskDeleteCmd.MarkFlagRequired("subtaskid")
}
