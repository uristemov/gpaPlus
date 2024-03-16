package service

import (
	"context"
	"fmt"
	"github.com/uristemov/repeatPro/api"
	"github.com/uristemov/repeatPro/internal/entity"
)

func (m *Manager) CreateModule(ctx context.Context, req *api.CreateModuleRequest) (string, error) {
	return m.Repository.CreateModule(ctx, req)
}

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

func (m *Manager) GetAllModuleSteps(ctx context.Context, id string) ([]api.GetStepsResponse, error) {
	return m.Repository.GetAllModuleSteps(ctx, id)
}

func (m *Manager) GetAllModuleWithSteps(ctx context.Context, id string) ([]api.GetModuleWithStepsResponse, error) {

	var responses []api.GetModuleWithStepsResponse

	modules, err := m.Repository.GetAllCourseModules(ctx, id)
	if err != nil {
		return nil, err
	}

	fmt.Println("Course modules: ", modules)

	for _, i := range modules {

		steps, err := m.Repository.GetAllModuleSteps(ctx, i.Id.String())
		if err != nil {
			return nil, err
		}

		response := api.GetModuleWithStepsResponse{
			Name:              i.Name,
			GetStepsResponses: steps,
		}

		responses = append(responses, response)
	}

	return responses, nil
}
