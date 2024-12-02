package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "CLI для управления задачами",
	Long: "This is a task management app. Each task can have subtasks.\n" +
		"Tasks are sorted by Deadline. Completed tasks are saved to the history.\n" +
		"Unexecuted tasks are also saved, but in another file and stored there until completion.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Используйте --help для просмотра доступных команд и их флагов.")
	},
}

func init() {
	rootCmd.Flags().BoolP("help", "h", false, "Показать справку по команде")
}

func Execute() error {
	return rootCmd.Execute()
}
