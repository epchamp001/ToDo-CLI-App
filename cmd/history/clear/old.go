package clear

import (
	"CLITodoApp/internal/entity"
	"fmt"
	"github.com/spf13/cobra"
)

var ClearOldCmd = &cobra.Command{
	Use:   "old",
	Short: "Clear history entries older than 7 days",
	Long:  "Remove all history entries that are older than 7 days from the history file.",
	Run: func(cmd *cobra.Command, args []string) {
		history := entity.NewHistory(1)

		if err := history.LoadFromFile("internal/repository/history.json"); err != nil {
			fmt.Println("Error loading history:", err)
			return
		}

		history.ClearOldEntry()
		history.ShowHistory()
		if err := history.SaveToFile("internal/repository/history.json"); err != nil {
			fmt.Println("Error saving history:", err)
			return
		}
		fmt.Println("Old history entries have been cleared successfully.")
	},
}

func init() {
}
