package http

import (
	"github.com/uristemov/repeatPro/internal/service"
	"go.uber.org/zap"
)

type Handler struct {
	service service.Service
	logger  *zap.SugaredLogger
}

func New(service service.Service, logger *zap.SugaredLogger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}
