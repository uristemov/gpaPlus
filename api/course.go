package api

type UpdateCourseRequest struct {
	Name        string  `json:"name" db:"name"`
	ImagePath   string  `json:"image_path" db:"image_path"`
	Description string  `json:"description" db:"description"`
	Price       int     `json:"price" db:"price"`
	Rating      float32 `json:"rating" db:"rating"`
}

type CreateCourseRequest struct {
	Name        string `json:"name" db:"name"`
	UserId      string `json:"user_id" db:"user_id"`
	ImagePath   string `json:"image_path" db:"image_path"`
	Description string `json:"description" db:"description"`
	Price       int    `json:"price" db:"price"`
}
