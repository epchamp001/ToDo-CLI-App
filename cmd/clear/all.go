package clear

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var ClearAllCmd = &cobra.Command{
	Use:   "all",
	Short: "Clear all data files",
	Long:  "Delete all data from id.json, tasks.json, and history.json, effectively resetting the application.",
	Run: func(cmd *cobra.Command, args []string) {
		files := map[string]string{
			"internal/repository/id.json":      "{}",
			"internal/repository/tasks.json":   "[]",
			"internal/repository/history.json": "{}",
		}

		for file, defaultContent := range files {
			if err := clearFile(file, defaultContent); err != nil {
				fmt.Printf("Error clearing file %s: %v\n", file, err)
				return
			}
		}
		fmt.Println("All data files have been successfully cleared.")
	},
}

func init() {
}

func clearFile(filePath, defaultContent string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to clear file: %w", err)
	}
	defer f.Close()

	_, err = f.WriteString(defaultContent)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}
	return nil
}
