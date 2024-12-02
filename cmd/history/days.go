package history

import (
	"CLITodoApp/internal/entity"
	"fmt"
	"github.com/spf13/cobra"
)

var days int

var DaysCmd = &cobra.Command{
	Use:   "days",
	Short: "Show history for the last N days",
	Long:  "Show the history of tasks that were completed or marked unnecessary in the last N days.",
	Run: func(cmd *cobra.Command, args []string) {
		history := entity.NewHistory(1)

		if err := history.LoadFromFile("internal/repository/history.json"); err != nil {
			fmt.Println("Error loading history:", err)
			return
		}

		filteredHistory, err := history.GetEntriesForLastDays(days)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		filteredHistory.ShowHistory()
	},
}

func init() {
	DaysCmd.Flags().IntVarP(&days, "days", "d", 0, "Number of days to filter history")
	DaysCmd.MarkFlagRequired("days")
}
