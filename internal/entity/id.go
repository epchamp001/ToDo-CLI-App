package entity

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type EntityID struct {
	TaskID         int `json:"task_id"`
	HistoryEntryID int `json:"history_entry_id"`
	HistoryID      int `json:"history_id"`
	sync.Mutex
}

func NewEntityID() *EntityID {
	return &EntityID{}
}

func (e *EntityID) GenerateTaskID() int {
	e.Lock()
	defer e.Unlock()
	e.TaskID++
	return e.TaskID
}

func (e *EntityID) GenerateHistoryEntryID() int {
	e.Lock()
	defer e.Unlock()
	e.HistoryEntryID++
	return e.HistoryEntryID
}

func (e *EntityID) GenerateHistoryID() int {
	e.Lock()
	defer e.Unlock()
	e.HistoryID++
	return e.HistoryID
}

func (e *EntityID) GenerateNewIDs() (taskID, historyEntryID, historyID int) {
	taskID = e.GenerateTaskID()
	historyEntryID = e.GenerateHistoryEntryID()
	historyID = e.GenerateHistoryID()
	return
}

func (e *EntityID) SaveToFile(path string) error {
	e.Lock()
	defer e.Unlock()
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to save ID file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(e); err != nil {
		return fmt.Errorf("failed to encode ID data: %w", err)
	}

	return nil
}

func (e *EntityID) LoadFromFile(path string) error {
	e.Lock()
	defer e.Unlock()

	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("couldn't find the file %v: %w", path, err)
		}
		return fmt.Errorf("failed to open ID file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(e); err != nil {
		return fmt.Errorf("failed to decode ID data: %w", err)
	}
	return nil
}

var ID = NewEntityID()
