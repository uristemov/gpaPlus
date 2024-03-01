package entity

import (
	"github.com/google/uuid"
	"time"
)

type University struct {
	Id        uuid.UUID `json:"id"  db:"id"`
	Name      string    `json:"name" binding:"required" db:"name"`
	ImagePath string    `json:"image_path" db:"image_path"`
	CreatedAt time.Time
}
