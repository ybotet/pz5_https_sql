package models

import "time"

type Task struct {
    ID          string     `json:"id"`
    Title       string     `json:"title"`
    Description string     `json:"description,omitempty"`
    Done        bool       `json:"done"`
    CreatedAt   time.Time  `json:"created_at"`
    DueDate     *time.Time `json:"due_date,omitempty"`
}