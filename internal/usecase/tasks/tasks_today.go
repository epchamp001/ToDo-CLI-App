package tasks

import (
	"CLITodoApp/internal/entity"
	"time"
)

func GetTasksWithTodayDeadline(tasks *[]*entity.Task) []*entity.Task {
	today := time.Now().Truncate(24 * time.Hour)

	var todayTasks []*entity.Task
	for _, task := range *tasks {
		if task.Deadline.Truncate(24 * time.Hour).Equal(today) {
			todayTasks = append(todayTasks, task)
		}
	}

	return todayTasks
}
