package service

import (
	"context"
	"fmt"
	"github.com/uristemov/repeatPro/api"
)

func (m *Manager) CreateStep(ctx context.Context, req *api.Step) (string, error) {
	switch (*req).(type) {
	case api.CreateVideoRequest:
		video, ok := (*req).(*api.CreateVideoRequest)
		if !ok {
			return "", fmt.Errorf("convert api.Step into api.CreateVideoRequest error")
		}
		return m.Repository.CreateVideo(ctx, video)
	}
	return "", fmt.Errorf("unsupported step type: %T", *req)
}
