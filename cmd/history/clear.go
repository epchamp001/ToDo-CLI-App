package history

import (
	cl "CLITodoApp/cmd/history/clear"
	"CLITodoApp/internal/entity"
	"fmt"
	"github.com/spf13/cobra"
)

var ClearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear the entire history",
	Long:  "Clear the entire history of completed or unnecessary tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		history := entity.NewHistory(1)

		if err := history.LoadFromFile("internal/repository/history.json"); err != nil {
			fmt.Println("Error loading history:", err)
			return
		}

		history.ClearAll()

		if err := history.SaveToFile("internal/repository/history.json"); err != nil {
			fmt.Println("Error saving cleared history:", err)
			return
		}

		fmt.Println("History has been successfully cleared.")
	},
}

func init() {
	ClearCmd.AddCommand(cl.ClearOldCmd)
}
