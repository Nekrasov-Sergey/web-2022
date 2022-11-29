package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"main/internal/app/config"
	"strconv"
)

const servicePrefix = "awesome_service."

type Client struct {
	cfg    config.RedisConfig
	client *redis.Client
}

func New(ctx context.Context, cfg config.RedisConfig) (*Client, error) {
	client := &Client{}

	client.cfg = cfg

	redisClient := redis.NewClient(&redis.Options{
		Username:    cfg.User,
		Addr:        cfg.Host + ":" + strconv.Itoa(cfg.Port),
		DB:          0,
		DialTimeout: cfg.DialTimeout,
		ReadTimeout: cfg.ReadTimeout,
	})

	client.client = redisClient

	if _, err := redisClient.Ping(ctx).Result(); err != nil {
		return nil, fmt.Errorf("cant ping redis: %w", err)
	}

	return client, nil
}

func (c *Client) Close() error {
	return c.client.Close()
}
