package api

type UpdateModuleRequest struct {
	Name string `json:"name" db:"name"`
}
