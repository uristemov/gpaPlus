package entity

import (
	"github.com/google/uuid"
	"time"
)

type Text struct {
	Id          uuid.UUID `json:"id"  db:"id"`
	Description string    `json:"description" binding:"required" db:"description"`
	Name        string    `json:"name" binding:"required" db:"name"`
	ModuleId    string    `json:"module_id" binding:"required" db:"module_id"`
	//LogoImage   string    `json:"logo_image" db:"logo_image"`
	//VideoPath string `json:"video_path" db:"video_path"`
	CreatedAt time.Time
}
