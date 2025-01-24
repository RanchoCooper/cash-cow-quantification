package repository

import (
	"context"

	mysql2 "quants/adapter/repository/mysql"
	redis2 "quants/adapter/repository/redis"
	"quants/config"
	"quants/util/logger"
)

var (
	Clients     = &client{}
	HealthCheck *redis2.HealthCheck
	Example     *mysql2.Example
	User        *mysql2.User
	Trade       *mysql2.Trade
)

type client struct {
	MySQL mysql2.IMySQL
	Redis redis2.IRedis
}

func (c *client) close(ctx context.Context) {
	if c.MySQL != nil {
		c.MySQL.Close(ctx)
	}
	if c.Redis != nil {
		c.Redis.Close(ctx)
	}
}

type Option func(*client)

func WithMySQL() Option {
	return func(c *client) {
		if c.MySQL == nil {
			if config.Config.MySQL != nil {
				c.MySQL = mysql2.NewMySQLClient()
			} else {
				panic("init repository with empty MySQL config")
			}
		}
		// inject repository
		if Example == nil {
			Example = mysql2.NewExample(Clients.MySQL)
		}
		if User == nil {
			User = mysql2.NewUser(Clients.MySQL)
		}
		if Trade == nil {
			Trade = mysql2.NewTrade(Clients.MySQL)
		}
	}
}

func WithRedis() Option {
	return func(c *client) {
		if c.Redis == nil {
			if config.Config.Redis != nil {
				c.Redis = redis2.NewRedisClient()
			} else {
				panic("init repository with empty Redis config")
			}
		}
		if HealthCheck == nil {
			HealthCheck = redis2.NewHealthCheck(Clients.Redis)
		}
	}
}

func Init(opts ...Option) {
	for _, opt := range opts {
		opt(Clients)
	}
	logger.Log.Info(context.Background(), "repository init successfully")
}

func Close(ctx context.Context) {
	Clients.close(ctx)
	logger.Log.Info(ctx, "repository is closed.")
}
