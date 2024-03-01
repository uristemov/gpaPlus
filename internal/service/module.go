package service

import (
	"context"
	"github.com/uristemov/repeatPro/api"
	"github.com/uristemov/repeatPro/internal/entity"
)

func (m *Manager) GetAllCourseModules(ctx context.Context, courseId string) ([]entity.Module, error) {
	return m.Repository.GetAllCourseModules(ctx, courseId)
}

func (m *Manager) GetModuleById(ctx context.Context, id string) (*entity.Module, error) {
	course, err := m.Repository.GetModuleById(ctx, id)
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (m *Manager) DeleteModuleById(ctx context.Context, id string) error {
	return m.Repository.DeleteModuleById(ctx, id)
}

func (m *Manager) UpdateModuleById(ctx context.Context, req *api.UpdateModuleRequest, id string) error {
	return m.Repository.UpdateModuleById(ctx, req, id)
}
