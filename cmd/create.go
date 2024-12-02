package cmd

import (
	"CLITodoApp/cmd/create"
	"CLITodoApp/internal/entity"
	tasks2 "CLITodoApp/internal/usecase/tasks"
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

var (
	description string
	deadline    string
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new task",
	Long:  "Create a new task with a description and deadline.",
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

		parsedDeadline, err := time.Parse("02-01-2006", deadline)
		if err != nil {
			fmt.Println("Error parsing deadline:", err)
			return
		}

		err = tasks2.CreateTask(description, parsedDeadline, &tasks)
		if err != nil {
			fmt.Println("Error creating task:", err)
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
	createCmd.Flags().StringVarP(&description, "description", "d", "", "Description of the task")
	createCmd.Flags().StringVarP(&deadline, "deadline", "l", "", "Deadline for the task (DD-MM-YYYY)")

	// указываем, что флаги обязательны
	createCmd.MarkFlagRequired("description")
	createCmd.MarkFlagRequired("deadline")

	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(create.SubtaskCmd)
}
