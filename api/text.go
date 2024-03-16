package api

type UpdateTextRequest struct {
	Description string `json:"description" db:"description"`
	Name        string `json:"name" binding:"required" db:"name"`
}

type CreateTextRequest struct {
	Description string `json:"description" db:"description"`
	ModuleId    string `json:"module_id" db:"module_id"`
	Name        string `json:"name" binding:"required" db:"name"`
}
