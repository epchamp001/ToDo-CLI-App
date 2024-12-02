package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "CLI для управления задачами",
	Long: "Это приложение для управления задачами. У каждой задачи могут быть подзадачи.\n" +
		"Задачи сортируются по Deadline. Выполненные задачи сохраняются в историю.\n" +
		"Невыполненные задачи также сохраняются, но в другой файл и там хранятся до выполнения.",
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
