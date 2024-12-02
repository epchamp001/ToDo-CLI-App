package create

import (
	"CLITodoApp/internal/entity"
	tasks2 "CLITodoApp/internal/usecase/tasks"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	parentID    int
	description string
)

var SubtaskCmd = &cobra.Command{
	Use:   "subtask",
	Short: "Create a new subtask",
	Long:  "Create a new subtask under a parent task with a specified ID, description, and deadline.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := tasks2.LoadTasksFromFile("internal/repository/tasks.json")
		if err != nil {
			fmt.Println("Error loading tasks from file:", err)
			return
		}

		if err := entity.ID.LoadFromFile("internal/repository/id.json"); err != nil {
			fmt.Println("Error loading ID file:", err)
			return
		}

		err = tasks2.CreateSubtask(description, parentID, &tasks)
		if err != nil {
			fmt.Println("Error creating subtask:", err)
			return
		}

		if err := tasks2.SaveTasksToFile("internal/repository/tasks.json", &tasks); err != nil {
			fmt.Println("Error saving tasks:", err)
			return
		}

		if err := entity.ID.SaveToFile("internal/repository/id.json"); err != nil {
			fmt.Println("Error saving ID file:", err)
			return
		}

		tasks2.ShowAllTasks(tasks)
	},
}

func init() {
	SubtaskCmd.Flags().IntVarP(&parentID, "parentid", "p", 0, "ID of the parent task")
	SubtaskCmd.Flags().StringVarP(&description, "description", "d", "", "Description of the subtask")

	SubtaskCmd.MarkFlagRequired("parentid")
	SubtaskCmd.MarkFlagRequired("description")
}
