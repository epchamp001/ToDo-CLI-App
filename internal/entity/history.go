package entity

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type History struct {
	ID      int             `json:"id"`
	History []*HistoryEntry `json:"history"`
}

func NewHistory(id int) *History {
	return &History{
		ID:      id,
		History: []*HistoryEntry{},
	}
}

func (h *History) AddEntry(entry *HistoryEntry) error {
	if entry == nil {
		return fmt.Errorf("invalid history entry")
	}

	for _, e := range h.History {
		if e.ID == entry.ID {
			return fmt.Errorf("entry with ID %d already exists", entry.ID)
		}
	}
	h.History = append(h.History, entry)
	return nil
}

func (h *History) RemoveEntry(id int) error {
	if len(h.History) == 0 {
		return fmt.Errorf("history is empty")
	}

	for i, e := range h.History {
		if e.ID == id {
			h.History = append(h.History[:i], h.History[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("entry with ID %d not found in history", id)
}

func (h *History) GetEntry(id int) (*HistoryEntry, error) {
	if len(h.History) == 0 {
		return nil, fmt.Errorf("history is empty")
	}

	for _, e := range h.History {
		if e.ID == id {
			return e, nil
		}
	}

	return nil, fmt.Errorf("entry with ID %d not found in history", id)
}

func (h *History) GetAllEntries() []*HistoryEntry {
	return h.History
}

func (h *History) ClearAll() {
	if len(h.History) != 0 {
		h.History = []*HistoryEntry{}
	}
}

func (h *History) ClearOldEntry() {
	if len(h.History) == 0 {
		return
	}

	newHistory := make([]*HistoryEntry, 0)
	for _, e := range h.History {
		if !e.IsOrderThanAWeek() {
			newHistory = append(newHistory, e)
		}
	}
	h.History = newHistory
}

func (h *History) GetEntriesForLastDays(days int) ([]*HistoryEntry, error) {
	if days <= 0 {
		return nil, fmt.Errorf("days must be greater than zero")
	}

	if len(h.History) == 0 {
		return nil, fmt.Errorf("history is empty")
	}

	// Время, которое было days дней назад
	t := time.Now().AddDate(0, 0, -days)

	result := make([]*HistoryEntry, 0)
	for _, e := range h.History {
		if e.Timestamp.After(t) || e.Timestamp.Equal(t) {
			result = append(result, e)
		}
	}

	return result, nil
}

func (h *History) GetEntriesForDate(date time.Time) ([]*HistoryEntry, error) {
	startOfDay := date.Truncate(24 * time.Hour)
	endOfDay := startOfDay.Add(24 * time.Hour)

	result := make([]*HistoryEntry, 0)

	for _, e := range h.History {
		if !e.Timestamp.Before(startOfDay) && e.Timestamp.Before(endOfDay) {
			result = append(result, e)
		}
	}

	return result, nil
}

func (h *History) SaveToFile(filename string) error {
	data, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal history: %w", err)
	}

	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("failed to save history to file: %w", err)
	}

	return nil
}

func (h *History) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	if err := json.Unmarshal(data, h); err != nil {
		return fmt.Errorf("failed to unmarshal history: %w", err)
	}

	return nil
}
