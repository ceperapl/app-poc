package models

import (
	"fmt"
	"time"
)

// Task entity
type Task struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (t Task) String() string {
	return fmt.Sprintf("Task{ID: %s, Description: %s, CreatedAt: %s, UpdatedAt: %s}", t.ID, t.Description, t.CreatedAt.Format("Mon Jan _2 15:04:05 2006"), t.UpdatedAt.Format("Mon Jan _2 15:04:05 2006"))
}
