package entity

import (
	"github.com/google/uuid"
	"time"
)

type Request struct {
	Id        uuid.UUID `json:"id" db:"id"`
	UserId    string    `json:"user_id" db:"user_id"`
	CourseId  string    `json:"course_id" db:"course_id"`
	IsActive  bool      `json:"is_active" db:"active"`
	CreatedAt time.Time
}
