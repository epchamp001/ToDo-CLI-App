package tasks

import (
	"CLITodoApp/internal/entity"
	"fmt"
)

func ReplaceTasks(taskID1, taskID2 int, tasks *[]*entity.Task) error {
	if len(*tasks) == 0 {
		return fmt.Errorf("task list is empty")
	}

	task1, task2, err := findTasksByID(taskID1, taskID2, *tasks)
	if err != nil {
		return err
	}

	swapTasks(task1, task2)
	SortTasksByDeadline(tasks)

	return nil
}

func findTasksByID(taskID1, taskID2 int, tasks []*entity.Task) (*entity.Task, *entity.Task, error) {
	var task1, task2 *entity.Task
	for _, t := range tasks {
		if t.ID == taskID1 {
			task1 = t
		}
		if t.ID == taskID2 {
			task2 = t
		}
	}

	if task1 == nil && task2 == nil {
		return nil, nil, fmt.Errorf("task with ID %d and task with ID %d not found", taskID1, taskID2)
	} else if task1 == nil {
		return nil, nil, fmt.Errorf("task with ID %d not found", taskID1)
	} else if task2 == nil {
		return nil, nil, fmt.Errorf("task with ID %d not found", taskID2)
	}

	return task1, task2, nil
}

func swapTasks(task1, task2 *entity.Task) {
	task1Deadline := task1.GetDeadline()
	task2Deadline := task2.GetDeadline()

	task1.UpdateDeadline(task2Deadline)
	task2.UpdateDeadline(task1Deadline)
}
