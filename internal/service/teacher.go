package service

import (
	"context"
	"fmt"
	"github.com/uristemov/repeatPro/internal/entity"
)

func (m *Manager) GetCoursesAndTeachers(ctx context.Context) ([]entity.Course, []entity.User, error) {

	courses, err := m.GetAllCourses(ctx)
	if err != nil {
		fmt.Println("Course error: ", err)
		return nil, nil, err
	}

	teachers, err := m.GetAllTeachers(ctx)
	if err != nil {
		fmt.Println("Teacher error: ", err)
		return nil, nil, err
	}

	return courses, teachers, nil
}

func (m *Manager) GetAllTeachers(ctx context.Context) ([]entity.User, error) {
	return m.Repository.GetAllTeachers(ctx)
}

func (m *Manager) GetTeacherById(ctx context.Context, id string) (*entity.User, error) {
	teacher, err := m.Repository.GetTeacherById(ctx, id)
	if err != nil {
		return nil, err
	}

	return teacher, nil
}
