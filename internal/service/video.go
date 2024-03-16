package service

import (
	"context"
	"github.com/uristemov/repeatPro/api"
	"github.com/uristemov/repeatPro/internal/entity"
)

func (m *Manager) CreateVideo(ctx context.Context, req *api.CreateVideoRequest) (string, error) {
	return m.Repository.CreateVideo(ctx, req)
}

func (m *Manager) GetAllVideos(ctx context.Context) ([]entity.Video, error) {
	return m.Repository.GetAllVideos(ctx)
}

func (m *Manager) GetVideoById(ctx context.Context, id string) (*entity.Video, error) {
	video, err := m.Repository.GetVideoById(ctx, id)
	if err != nil {
		return nil, err
	}

	return video, nil
}

func (m *Manager) DeleteVideoById(ctx context.Context, id string) error {
	return m.Repository.DeleteVideoById(ctx, id)
}

func (m *Manager) UpdateVideoById(ctx context.Context, req *api.UpdateVideoRequest, id string) error {
	return m.Repository.UpdateVideoById(ctx, req, id)
}
