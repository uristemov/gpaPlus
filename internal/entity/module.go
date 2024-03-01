package entity

import (
	"github.com/google/uuid"
)

type Module struct {
	Id       uuid.UUID `json:"id"  db:"id"`
	CourseId uuid.UUID `json:"course_id"  db:"course_id"`
	Name     string    `json:"name" db:"name"`
}
