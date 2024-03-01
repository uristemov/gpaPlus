package entity

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id           uuid.UUID `json:"id"  db:"id"`
	FirstName    string    `json:"firstname" binding:"required" db:"first_name"`
	LastName     string    `json:"lastname" binding:"required" db:"last_name"`
	ImagePath    string    `json:"image_path" db:"image_path"`
	Phone        string    `json:"phone" db:"phone"`
	Password     string    `json:"password" binding:"required" db:"password"`
	UniversityId string    `json:"university_id"  db:"university_id"`
	RoleId       int64     `json:"role_id"  db:"role_id"`
	Verified     bool      `json:"verified"  db:"verified"`
	Email        string    `json:"email" binding:"required" db:"email"`
	CreatedAt    time.Time
}
