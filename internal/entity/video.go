package entity

import (
	"github.com/google/uuid"
	"time"
)

type Video struct {
	Id          uuid.UUID `json:"id"  db:"id"`
	Description string    `json:"description" db:"description"`
	LogoImage   string    `json:"logo_image" db:"logo_image"`
	VideoPath   string    `json:"video_path" db:"video_path"`
	CreatedAt   time.Time
}
