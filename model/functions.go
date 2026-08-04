package model

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

func (tl *TaskList) AddTask(task Task) {
	tl.NewID++
	task.ID = strconv.FormatUint(tl.NewID, 10)
	tl.mux.Lock()
	defer tl.mux.Unlock()
	tl.Tasks = append(tl.Tasks, task)
}

func (tl *TaskList) SaveToFile(path string) error {
	tl.mux.RLock()
	defer tl.mux.RUnlock()

	data, err := json.MarshalIndent(tl, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
func (tl *TaskList) LoadFromFile(path string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	tl.mux.Lock()
	defer tl.mux.Unlock()

	decoder := json.NewDecoder(file)

	if err := decoder.Decode(tl); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (t Task) String() string {
	var status string
	if t.Done == true {
		status = "done"
	} else {
		status = "running"
	}
	return fmt.Sprintf("id:%s\n- title: %s\n- Description: %s\n- status:%s", t.ID, t.Title, t.Description, status)
}
