package service

import (
	"github.com/Qwepo/InCryipt/internal"
	"github.com/Qwepo/InCryipt/internal/repository"
	"github.com/redis/go-redis/v9"
)

type Service interface {
	sAuth
}

type service struct {
	repo  repository.Repository
	redis *redis.Client
	conf  *internal.Config
}

func NewService(repo repository.Repository, redis *redis.Client, config *internal.Config) Service {
	return &service{repo: repo, redis: redis, conf: config}
}
