package api

type UpdateModuleRequest struct {
	Name string `json:"name" db:"name"`
}

type CreateModuleRequest struct {
	CourseId string `json:"course_id"  db:"course_id"`
	Name     string `json:"name" db:"name"`
}

type GetModuleWithStepsResponse struct {
	Name              string             `json:"name" db:"name"`
	GetStepsResponses []GetStepsResponse `json:"steps"`
}

type GetStepsResponse struct {
	Id   string `json:"id"  db:"id"`
	Name string `json:"name" db:"name"`
	Url  string `json:"url"`
}
