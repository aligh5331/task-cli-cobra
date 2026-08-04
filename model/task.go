package model

import (
	"sync"
)

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
	NewID uint64 `json:"new_id"`
	mux   sync.RWMutex
}
