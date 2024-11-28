package tasks

import (
	"CLITodoApp/internal/entity"
	"sort"
)

// SortTasksByDeadline сортирует задачи по следующему приоритету:
// 1. По сроку (Deadline)
// 2. Если сроки равны, то по времени создания (CreatedAt)
// 3. Если оба эти поля равны, то по ID
func SortTasksByDeadline(t *[]*entity.Task) {
	sort.Sort(ByDeadline(*t))
	return
}

// ByDeadline реализует интерфейс sort.Interface для сортировки по нескольким критериям
type ByDeadline []*entity.Task

func (b ByDeadline) Len() int {
	return len(b)
}

// Less определяет порядок сортировки
func (b ByDeadline) Less(i, j int) bool {
	if b[i].Deadline != b[j].Deadline {
		return b[i].Deadline.Before(b[j].Deadline)
	}

	if b[i].CreatedAt != b[j].CreatedAt {
		return b[i].CreatedAt.Before(b[j].CreatedAt)
	}

	return b[i].ID < b[j].ID
}

func (b ByDeadline) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
