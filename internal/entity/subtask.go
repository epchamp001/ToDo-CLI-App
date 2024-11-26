package entity

type Subtask struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      Status `json:"status"`
}
