package api

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	FirstName    string `json:"firstname" db:"firstname"`
	LastName     string `json:"lastname" db:"lastname"`
	ImagePath    string `json:"image_path" db:"image_path"`
	Phone        string `json:"phone" db:"phone"`
	Password     string `json:"password" db:"password"`
	UniversityId string `json:"university_id"  db:"university_id"`
	RoleId       int64  `json:"role_id"  db:"role_id"`
	Email        string `json:"email" db:"email" gorm:"unique"`
	Verified     bool   `json:"verified"  db:"verified"`
}

type UpgradeUserRequest struct {
	UniversityId string `json:"university_id"  db:"university_id"`
	RoleId       int64  `json:"role_id"  db:"role_id"`
	Verified     bool   `json:"verified"  db:"verified"`
}
