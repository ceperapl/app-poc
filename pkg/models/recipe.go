package models

import (
	"fmt"
	"time"
)

// Recipe entity
type Recipe struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
	Steps       []string     `json:"steps"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
}

type Ingredient struct {
	Name        string  `json:"name"`
	Amount      float32 `json:"amount"`
	Measurement string  `json:"measurement"`
}

func (t Recipe) String() string {
	return fmt.Sprintf("Recipe{ID: %s, Name: %s, CreatedAt: %s, UpdatedAt: %s}", t.ID, t.Name, t.CreatedAt.Format("Mon Jan _2 15:04:05 2006"), t.UpdatedAt.Format("Mon Jan _2 15:04:05 2006"))
}
