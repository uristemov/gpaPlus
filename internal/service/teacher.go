package service

import (
	"context"
	"github.com/uristemov/repeatPro/internal/entity"
)

func (m *Manager) GetAllTeachers(ctx context.Context) ([]entity.Teacher, error) {
	return m.Repository.GetAllTeachers(ctx)
}

func (m *Manager) GetTeacherById(ctx context.Context, id string) (*entity.Teacher, error) {
	teacher, err := m.Repository.GetTeacherById(ctx, id)
	if err != nil {
		return nil, err
	}

	return teacher, nil
}
