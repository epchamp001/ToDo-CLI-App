package cmd

import (
	clear2 "CLITodoApp/cmd/clear"
	"CLITodoApp/internal/entity"
	history2 "CLITodoApp/internal/usecase/history"
	tasks2 "CLITodoApp/internal/usecase/tasks"
	"fmt"
	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear completed or unnecessary tasks",
	Long:  "Remove completed or unnecessary tasks from the current list and save them to history.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := tasks2.LoadTasksFromFile("internal/repository/tasks.json")
		if err != nil {
			fmt.Println("Error loading tasks from file:", err)
			return
		}

		activeTasks, inactiveTasks := tasks2.SplitTasksByStatus(&tasks)

		if err := tasks2.SaveTasksToFile("internal/repository/tasks.json", &activeTasks); err != nil {
			fmt.Println("Error saving active tasks:", err)
			return
		}

		history := entity.NewHistory(1)

		if err := history.LoadFromFile("internal/repository/history.json"); err != nil {
			fmt.Println("Error loading history:", err)
			return
		}

		if err := entity.ID.LoadFromFile("internal/repository/id.json"); err != nil {
			fmt.Println("Error loading ID file:", err)
			return
		}

		for _, task := range inactiveTasks {
			if err := history2.CreateEntryInHistory(task, history); err != nil {
				fmt.Printf("Error adding task with ID %d to history: %w\n", task.ID, err)
			}
		}

		if err := history.SaveToFile("internal/repository/history.json"); err != nil {
			fmt.Println("Error saving history:", err)
			return
		}

		if err := entity.ID.SaveToFile("internal/repository/id.json"); err != nil {
			fmt.Println("Error saving ID file:", err)
			return
		}

		tasks2.ShowAllTasks(activeTasks)
	},
}

func init() {
	clearCmd.AddCommand(clear2.ClearAllCmd)
	rootCmd.AddCommand(clearCmd)
}
