package history

import (
	"CLITodoApp/internal/entity"
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

var date string

var DateCmd = &cobra.Command{
	Use:   "date",
	Short: "Show history for a specific date",
	Long:  "Show the history of tasks completed or marked unnecessary on a specific date.",
	Run: func(cmd *cobra.Command, args []string) {
		history := entity.NewHistory(1)

		if err := history.LoadFromFile("internal/repository/history.json"); err != nil {
			fmt.Println("Error loading history:", err)
			return
		}

		parsedDate, err := time.Parse("02-01-2006", date)
		if err != nil {
			fmt.Println("Invalid date format. Use DD-MM-YYYY.")
			return
		}

		filteredHistory, err := history.GetEntriesForDate(parsedDate)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		filteredHistory.ShowHistory()
	},
}

func init() {
	DateCmd.Flags().StringVarP(&date, "date", "d", "", "Date to filter history (format: DD-MM-YYYY)")
	DateCmd.MarkFlagRequired("date")
}
