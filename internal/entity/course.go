package entity

import (
	"github.com/google/uuid"
	"time"
)

type Course struct {
	Id          uuid.UUID `json:"course_id" db:"id"`
	Name        string    `json:"name" db:"name"`
	ImagePath   string    `json:"image_path" db:"image_path"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time
}
