package api

type UpdateImageRequest struct {
	Description string `json:"description" db:"description"`
	//LogoImage   string `json:"logo_image" db:"logo_image"`
	ImagePath string `json:"image_path" db:"image_path"`
	Name      string `json:"name" db:"name"`
}

type CreateImageRequest struct {
	Description string `json:"description" db:"description"`
	ModuleId    string `json:"module_id" db:"module_id"`
	ImagePath   string `json:"image_path" binding:"required" db:"image_path"`
	Name        string `json:"name" binding:"required" db:"name"`
}
