package entity

import (
	"github.com/google/uuid"
	"time"
)

type Image struct {
	Id          uuid.UUID `json:"id"  db:"id"`
	Description string    `json:"description" db:"description"`
	Name        string    `json:"name" binding:"required" db:"name"`
	ModuleId    string    `json:"module_id" binding:"required" db:"module_id"`
	ImagePath   string    `json:"image_path" binding:"required" db:"image_path"`
	CreatedAt   time.Time
}
