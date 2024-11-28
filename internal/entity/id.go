package entity

import (
	"sync"
)

type EntityID struct {
	TaskID         int
	HistoryEntryID int
	HistoryID      int
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

var ID = NewEntityID()
