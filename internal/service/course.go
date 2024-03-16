package service

import (
	"context"
	"github.com/uristemov/repeatPro/api"
	"github.com/uristemov/repeatPro/internal/entity"
)

func (m *Manager) CreateCourse(ctx context.Context, req *api.CreateCourseRequest) (string, error) {
	return m.Repository.CreateCourse(ctx, req)
}

func (m *Manager) GetAllCourses(ctx context.Context) ([]entity.Course, error) {
	return m.Repository.GetAllCourses(ctx)
}

func (m *Manager) GetCourseById(ctx context.Context, id string) (*entity.Course, error) {
	course, err := m.Repository.GetCourseById(ctx, id)
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (m *Manager) DeleteCourseById(ctx context.Context, id string) error {
	return m.Repository.DeleteCourseById(ctx, id)
}

func (m *Manager) UpdateCourseById(ctx context.Context, req *api.UpdateCourseRequest, id string) error {
	return m.Repository.UpdateCourseById(ctx, req, id)
}

func (m *Manager) GetAllTeacherCourses(ctx context.Context, id string) ([]entity.Course, error) {
	return m.Repository.GetAllTeacherCourses(ctx, id)
}

func (m *Manager) AddStudentToCourse(ctx context.Context, userId, courseId string) error {
	return m.Repository.AddStudentToCourse(ctx, userId, courseId)
}

func (m *Manager) GetAllCourseStudents(ctx context.Context, userId, courseId string) ([]api.GetStudentsResponse, error) {
	_, err := m.Repository.GetCourseById(ctx, courseId)
	if err != nil {
		return nil, err
	}

	return m.Repository.GetAllCourseStudents(ctx, userId, courseId)
}
