package applicator

import (
	"context"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/uristemov/repeatPro/internal/cache"
	"github.com/uristemov/repeatPro/internal/config"
	"github.com/uristemov/repeatPro/internal/controller/http"
	"github.com/uristemov/repeatPro/internal/repository"
	"github.com/uristemov/repeatPro/internal/service"
	redisCache "github.com/uristemov/repeatPro/pkg/cache"
	"github.com/uristemov/repeatPro/pkg/front"
	"github.com/uristemov/repeatPro/pkg/go_admin"
	"github.com/uristemov/repeatPro/pkg/go_admin/tables"
	"github.com/uristemov/repeatPro/pkg/httpserver"
	"github.com/uristemov/repeatPro/pkg/jwt_token"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func Run(logger *zap.SugaredLogger, cfg *config.Config) {

	ctx, cancel := context.WithCancel(context.TODO())
	_ = ctx

	db, err := repository.New(
		repository.WithHost(cfg.Database.Host),
		repository.WithUsername(cfg.Database.Username),
		repository.WithPort(cfg.Database.Port),
		repository.WithDBName(cfg.Database.DBName),
		repository.WithPassword(cfg.Database.Password),
	)
	if err != nil {
		logger.Panicf("failed connect to DB at '%s:%d': %v", cfg.Database.Host, cfg.Database.Port, err)
	}

	logger.Infof("Successfully connected to db at %d port", cfg.Database.Port)

	token := jwt_token.New(os.Getenv("TOKEN_SECRET_KEY"))

	redisClient, err := redisCache.NewRedisClient(cfg)
	if err != nil {
		logger.Infof("failed connect to DB at '%s:%d': %v", err)
	}

	appCache := cache.NewCache(redisClient, cfg.Redis.ExpirationTime)

	cacher, err := cache.NewAppCache(cache.WithUserCache(appCache), cache.WithTokenCache(appCache))
	if err != nil {
		logger.Errorf("failed create necessary caches err: %w", err)
	}

	service := service.New(db, cfg, token, cacher, logger)
	handler := http.New(service, logger)

	router := handler.InitRouter()

	router = front.InitFront(router)

	goAdmin := tables.New(service)

	eng := go_admin.InitGoAdmin(router, cfg.Database, logger, goAdmin)

	server, err := httpserver.NewServer(cfg.HttpServer.Port, cfg.HttpServer.ShutdownTimeout, logger, router)
	if err != nil {
		logger.Panicf("failed to create, initate server error: %v", err)
	}

	server.Run()
	logger.Infof("HTTP server is running at %d port", cfg.HttpServer.Port)

	defer func() {
		if err := server.Stop(); err != nil {
			logger.Panicf("failed close server err: %v", err)
		}
		logger.Info("server closed")
	}()

	gracefulShutdown(cancel, eng)
}

func gracefulShutdown(cancel context.CancelFunc, eng *engine.Engine) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	<-ch
	signal.Stop(ch)
	eng.PostgresqlConnection().Close()
	cancel()
}
