package entity

import (
	"time"

	"github.com/google/uuid"
)

type Teacher struct {
	Id          uuid.UUID `json:"teacher_id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Phone       string    `json:"phone" db:"phone"`
	ImagePath   string    `json:"image_path" db:"image_path"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time
}
