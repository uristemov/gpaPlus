package service

import (
	"context"
	"github.com/uristemov/repeatPro/api"
	"github.com/uristemov/repeatPro/internal/entity"
)

func (m *Manager) CreateText(ctx context.Context, req *api.CreateTextRequest) (string, error) {
	return m.Repository.CreateText(ctx, req)
}

func (m *Manager) GetAllTexts(ctx context.Context) ([]entity.Text, error) {
	return m.Repository.GetAllTexts(ctx)
}

func (m *Manager) GetTextById(ctx context.Context, id string) (*entity.Text, error) {
	video, err := m.Repository.GetTextById(ctx, id)
	if err != nil {
		return nil, err
	}

	return video, nil
}

func (m *Manager) DeleteTextById(ctx context.Context, id string) error {
	return m.Repository.DeleteTextById(ctx, id)
}

func (m *Manager) UpdateTextById(ctx context.Context, req *api.UpdateTextRequest, id string) error {
	return m.Repository.UpdateTextById(ctx, req, id)
}
