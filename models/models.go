package models

// Task represents a task entity.
type Task struct {
	TaskContent string `json:"task_content"`
	NumInList   int    `json:"num_in_list"`
}

