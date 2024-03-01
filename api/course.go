package api

type UpdateCourseRequest struct {
	Name        string `json:"name" db:"name"`
	ImagePath   string `json:"image_path" db:"image_path"`
	Description string `json:"description" db:"description"`
}
