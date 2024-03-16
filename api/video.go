package api

type UpdateVideoRequest struct {
	Description string `json:"description" db:"description"`
	//LogoImage   string `json:"logo_image" db:"logo_image"`
	VideoPath string `json:"video_path" db:"video_path"`
	Name      string `json:"name" binding:"required" db:"name"`
}

type CreateVideoRequest struct {
	Description string `json:"description" db:"description"`
	ModuleId    string `json:"module_id" db:"module_id"`
	VideoPath   string `json:"video_path" db:"video_path"`
	Name        string `json:"name" binding:"required" db:"name"`
}

func (v CreateVideoRequest) GetDescription() string {
	return v.Description
}

func (v CreateVideoRequest) GetName() string {
	return v.Name
}
