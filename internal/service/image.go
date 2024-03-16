package service

import (
	"context"
	"github.com/uristemov/repeatPro/api"
	"github.com/uristemov/repeatPro/internal/entity"
)

func (m *Manager) CreateImage(ctx context.Context, req *api.CreateImageRequest) (string, error) {
	return m.Repository.CreateImage(ctx, req)
}

func (m *Manager) GetAllImages(ctx context.Context) ([]entity.Image, error) {
	return m.Repository.GetAllImages(ctx)
}

func (m *Manager) GetImageById(ctx context.Context, id string) (*entity.Image, error) {
	return m.Repository.GetImageById(ctx, id)
}

func (m *Manager) DeleteImageById(ctx context.Context, id string) error {
	return m.Repository.DeleteImageById(ctx, id)
}

func (m *Manager) UpdateImageById(ctx context.Context, req *api.UpdateImageRequest, id string) error {
	return m.Repository.UpdateImageById(ctx, req, id)
}
