package app

import (
	"context"
	log "github.com/sirupsen/logrus"
	"main/internal/app/config"
	"main/internal/app/dsn"
	"main/internal/app/redis"
	"main/internal/app/repository"
)

type Application struct {
	config *config.Config
	repo   *repository.Repository
	redis  *redis.Client
}

func New(ctx context.Context) (*Application, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	repo, err := repository.New(dsn.FromEnv())
	if err != nil {
		return nil, err
	}

	redisClient, err := redis.New(ctx, cfg.Redis)
	if err != nil {
		return nil, err
	}

	return &Application{
		config: cfg,
		repo:   repo,
		redis:  redisClient,
	}, nil
}

func (a *Application) Run() error {
	log.Println("application start running")
	a.StartServer()
	log.Println("application shut down")

	return nil
}
