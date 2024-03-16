package api

type CreateRequest struct {
	UserId   string `json:"user_id" db:"user_id"`
	CourseId string `json:"course_id" db:"course_id"`
	IsActive bool   `json:"is_active" db:"active"`
}

type UpdateRequest struct {
	Accepted bool `json:"accepted" db:"accepted"`
}
