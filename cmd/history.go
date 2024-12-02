package cmd

import (
	"CLITodoApp/cmd/history"
	"CLITodoApp/internal/entity"
	"fmt"
	"github.com/spf13/cobra"
)

var days int

var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Show the history of tasks",
	Long:  "Show the entire history of completed and unnecessary tasks stored in history.json.",
	Run: func(cmd *cobra.Command, args []string) {
		// Если потребуются разный истории, то id для history подгружать из id.json
		// И тогда передавать в функцию создания NewHistory уже метод entity.ID.GenerateHistoryID()
		history := entity.NewHistory(1)

		if err := history.LoadFromFile("internal/repository/history.json"); err != nil {
			fmt.Println("Error loading history:", err)
			return
		}

		history.ShowHistory()
	},
}

func init() {
	historyCmd.AddCommand(history.DaysCmd)
	historyCmd.AddCommand(history.DateCmd)
	historyCmd.AddCommand(history.ClearCmd)
	historyCmd.AddCommand(history.TaskCmd)
	rootCmd.AddCommand(historyCmd)
}
