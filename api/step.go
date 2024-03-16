package api

type Step interface {
	GetDescription() string
	GetName() string
}

type Steps struct {
	Id   string `json:"id"  db:"id"`
	Name string `json:"name" db:"name"`
}
