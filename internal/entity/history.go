package entity

type History struct {
	ID      int            `json:"id"`
	History []HistoryEntry `json:"history"`
}
