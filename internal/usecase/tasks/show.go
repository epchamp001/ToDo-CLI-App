package tasks

import (
	"CLITodoApp/internal/entity"
	"fmt"
	"strings"
)

func ShowAllTasks(tasks []*entity.Task) {

	width := 80

	maxIDWidth := getMaxIDWidth(tasks)
	maxDescWidth := width - maxIDWidth - 20 // Учитываем место для дедлайна и статуса

	for i, task := range tasks {
		if i > 0 {
			fmt.Println()
			fmt.Println()
		}
		fmt.Println(formatTask(task, 0, maxIDWidth, maxDescWidth, width))
	}
}

func formatTask(task *entity.Task, level int, idWidth, descWidth, consoleWidth int) string {
	statusEmoji := getEmoji(task.Status)

	indent := strings.Repeat("    ", level)

	description := truncateString(task.Description, descWidth)
	formatted := fmt.Sprintf(
		"%s%s %-*s %s %s",
		indent,
		statusEmoji,
		descWidth, description,
		alignRight(task.Deadline.Format("2006-01-02"), consoleWidth-len(indent)-idWidth-descWidth-5),
		alignRight(fmt.Sprintf("(ID: %d)", task.ID), 20),
	)

	for i, subtask := range task.Subtask {
		if i > 0 {
			formatted += "\n"
		}
		formatted += "\n" + formatTask(subtask, level+1, idWidth, descWidth, consoleWidth)
	}

	return formatted
}

func truncateString(s string, maxWidth int) string {
	if len(s) > maxWidth {
		return s[:maxWidth-3] + "..."
	}
	return s
}

func alignRight(s string, width int) string {
	if len(s) < width {
		return strings.Repeat(" ", width-len(s)) + s
	}
	return s
}

func getMaxIDWidth(tasks []*entity.Task) int {
	maxWidth := 0
	for _, task := range tasks {
		width := len(fmt.Sprintf("(ID: %d)", task.ID))
		if width > maxWidth {
			maxWidth = width
		}
		for _, subtask := range task.Subtask {
			subWidth := len(fmt.Sprintf("(ID: %d)", subtask.ID))
			if subWidth > maxWidth {
				maxWidth = subWidth
			}
		}
	}
	return maxWidth
}

func getEmoji(status entity.Status) string {
	var statusEmoji string
	switch status {
	case entity.Done:
		statusEmoji = "✅"
	case entity.Active:
		statusEmoji = "⬜️"
	case entity.Unnecessary:
		statusEmoji = "❌"
	default:
		statusEmoji = "❓"
	}
	return statusEmoji
}

func ShowTask(task *entity.Task) {
	statusEmoji := getEmoji(task.Status)

	wrappedDescription := wrapText(task.Description, 80)
	fmt.Printf("ID: %d\n", task.ID)
	fmt.Printf("Created At: %s\n", task.CreatedAt.Format("02-01-2006 15:04"))
	fmt.Printf("Deadline: %s\n", task.Deadline.Format("02-01-2006 15:04"))
	fmt.Printf("Description: %s\n", wrappedDescription)
	fmt.Printf("Status: %s %s\n", statusEmoji, task.Status)
	fmt.Println("------------")

	if len(task.Subtask) > 0 {
		fmt.Println()
		fmt.Println("Subtasks:")
		for _, subtask := range task.Subtask {
			ShowTask(subtask)
		}
	}

}

func wrapText(text string, width int) string {
	var result []string
	for len(text) > width {
		cut := width
		if spaceIdx := strings.LastIndex(text[:cut], " "); spaceIdx != -1 {
			cut = spaceIdx
		}

		result = append(result, text[:cut])
		text = strings.TrimSpace(text[cut:])
	}
	result = append(result, text)
	return strings.Join(result, "\n")
}
