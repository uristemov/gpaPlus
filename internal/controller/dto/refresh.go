package dto

type RefreshInput struct {
	Token string `json:"refreshToken" binding:"required"`
}
