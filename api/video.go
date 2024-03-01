package api

type UpdateVideoRequest struct {
	Description string `json:"description" db:"description"`
	LogoImage   string `json:"logo_image" db:"logo_image"`
	VideoPath   string `json:"video_path" db:"video_path"`
}
