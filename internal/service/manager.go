package service

import (
	"github.com/uristemov/repeatPro/internal/cache"
	"github.com/uristemov/repeatPro/internal/config"
	"github.com/uristemov/repeatPro/internal/repository"
	"github.com/uristemov/repeatPro/pkg/jwt_token"
	"go.uber.org/zap"
)

type Manager struct {
	Repository repository.Repository
	Config     *config.Config
	Token      *jwt_token.JWTToken
	Cache      *cache.AppCache
	logger     *zap.SugaredLogger
}

func New(repository repository.Repository, config *config.Config, token *jwt_token.JWTToken, cache *cache.AppCache, logger *zap.SugaredLogger) *Manager {
	return &Manager{Repository: repository, Config: config, Token: token, Cache: cache, logger: logger}
}
