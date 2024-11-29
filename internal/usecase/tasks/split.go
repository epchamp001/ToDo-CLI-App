package tasks

import "CLITodoApp/internal/entity"

func SplitTasksByStatus(tasks *[]*entity.Task) ([]*entity.Task, []*entity.Task) {
	activeTask := make([]*entity.Task, 0)
	inactiveTask := make([]*entity.Task, 0)

	for _, task := range *tasks {
		if task.GetStatus() == entity.Active {
			activeTask = append(activeTask, task)
		} else {
			inactiveTask = append(inactiveTask, task)
		}
	}
	return activeTask, inactiveTask
}
