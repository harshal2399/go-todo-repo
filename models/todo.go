package models

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	Id          uuid.UUID `json:"id" gorm : "primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at" gorm:"type:timestamptz"`
}
