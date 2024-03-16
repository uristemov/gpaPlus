package service

import (
	"context"
	"github.com/uristemov/repeatPro/api"
	"github.com/uristemov/repeatPro/internal/entity"
)

func (m *Manager) GetAllTeacherRequests(ctx context.Context, id string) ([]entity.Request, error) {
	return m.Repository.GetAllTeacherRequests(ctx, id)
}

func (m *Manager) CreateRequest(ctx context.Context, req *api.CreateRequest) (string, error) {
	return m.Repository.CreateRequest(ctx, req)
}

func (m *Manager) UpdateRequestById(ctx context.Context, req *api.UpdateRequest, id string) error {

	request, err := m.Repository.GetRequestById(ctx, id)
	if err != nil {
		return err
	}

	if req.Accepted {
		err = m.Repository.AddStudentToCourse(ctx, request.UserId, request.CourseId)
		if err != nil {
			return err
		}
	}

	return m.Repository.DeleteRequestById(ctx, id)
}

func (m *Manager) DeleteRequestById(ctx context.Context, id string) error {
	return m.Repository.DeleteRequestById(ctx, id)
}

func (m *Manager) GetRequestById(ctx context.Context, id string) (*entity.Request, error) {
	return m.Repository.GetRequestById(ctx, id)
}
