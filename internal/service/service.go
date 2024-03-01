package service

import (
	"context"
	"github.com/uristemov/repeatPro/api"
	"github.com/uristemov/repeatPro/internal/entity"
)

type Service interface {
	CreateUser(ctx context.Context, u *entity.User) (string, error)
	UpdateUser(ctx context.Context, id string, req *api.UpdateUserRequest) error
	GetUserById(ctx context.Context, id string) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	DeleteUser(ctx context.Context, id string) error
	UpgradeUser(ctx context.Context, id string, req *api.UpdateUserRequest) error

	Login(ctx context.Context, email, password string) (string, string, error)
	VerifyToken(token string) (string, error)
	Refresh(ctx context.Context, refreshToken string) (string, string, error)

	GetAllTeachers(ctx context.Context) ([]entity.Teacher, error)
	GetTeacherById(ctx context.Context, id string) (*entity.Teacher, error)

	GetAllCourses(ctx context.Context) ([]entity.Course, error)
	GetCourseById(ctx context.Context, id string) (*entity.Course, error)
	DeleteCourseById(ctx context.Context, id string) error
	UpdateCourseById(ctx context.Context, req *api.UpdateCourseRequest, id string) error

	GetAllCourseModules(ctx context.Context, courseId string) ([]entity.Module, error)
	GetModuleById(ctx context.Context, id string) (*entity.Module, error)
	DeleteModuleById(ctx context.Context, id string) error
	UpdateModuleById(ctx context.Context, req *api.UpdateModuleRequest, id string) error

	GetAllVideos(ctx context.Context) ([]entity.Video, error)
	GetVideoById(ctx context.Context, id string) (*entity.Video, error)
	DeleteVideoById(ctx context.Context, id string) error
	UpdateVideoById(ctx context.Context, req *api.UpdateVideoRequest, id string) error
}
