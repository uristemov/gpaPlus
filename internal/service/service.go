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

	GetCoursesAndTeachers(ctx context.Context) ([]entity.Course, []entity.User, error)

	GetAllTeachers(ctx context.Context) ([]entity.User, error)
	GetTeacherById(ctx context.Context, id string) (*entity.User, error)
	GetAllCourseStudents(ctx context.Context, userId, courseId string) ([]api.GetStudentsResponse, error)

	GetAllTeacherRequests(ctx context.Context, id string) ([]entity.Request, error)
	CreateRequest(ctx context.Context, req *api.CreateRequest) (string, error)
	UpdateRequestById(ctx context.Context, req *api.UpdateRequest, id string) error
	GetRequestById(ctx context.Context, id string) (*entity.Request, error)

	CreateCourse(ctx context.Context, req *api.CreateCourseRequest) (string, error)
	GetAllCourses(ctx context.Context) ([]entity.Course, error)
	GetCourseById(ctx context.Context, id string) (*entity.Course, error)
	DeleteCourseById(ctx context.Context, id string) error
	UpdateCourseById(ctx context.Context, req *api.UpdateCourseRequest, id string) error
	GetAllTeacherCourses(ctx context.Context, id string) ([]entity.Course, error)
	AddStudentToCourse(ctx context.Context, userId, courseId string) error

	CreateModule(ctx context.Context, req *api.CreateModuleRequest) (string, error)
	GetAllCourseModules(ctx context.Context, courseId string) ([]entity.Module, error)
	GetModuleById(ctx context.Context, id string) (*entity.Module, error)
	DeleteModuleById(ctx context.Context, id string) error
	UpdateModuleById(ctx context.Context, req *api.UpdateModuleRequest, id string) error
	GetAllModuleSteps(ctx context.Context, id string) ([]api.GetStepsResponse, error)
	GetAllModuleWithSteps(ctx context.Context, id string) ([]api.GetModuleWithStepsResponse, error)

	CreateVideo(ctx context.Context, req *api.CreateVideoRequest) (string, error)
	GetAllVideos(ctx context.Context) ([]entity.Video, error)
	GetVideoById(ctx context.Context, id string) (*entity.Video, error)
	DeleteVideoById(ctx context.Context, id string) error
	UpdateVideoById(ctx context.Context, req *api.UpdateVideoRequest, id string) error

	CreateText(ctx context.Context, req *api.CreateTextRequest) (string, error)
	GetAllTexts(ctx context.Context) ([]entity.Text, error)
	GetTextById(ctx context.Context, id string) (*entity.Text, error)
	DeleteTextById(ctx context.Context, id string) error
	UpdateTextById(ctx context.Context, req *api.UpdateTextRequest, id string) error

	CreateImage(ctx context.Context, req *api.CreateImageRequest) (string, error)
	GetAllImages(ctx context.Context) ([]entity.Image, error)
	GetImageById(ctx context.Context, id string) (*entity.Image, error)
	DeleteImageById(ctx context.Context, id string) error
	UpdateImageById(ctx context.Context, req *api.UpdateImageRequest, id string) error
}
