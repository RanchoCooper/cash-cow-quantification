package repository

import (
	"context"
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"

	"quants/adapter/repository/mysql"
	"quants/adapter/repository/redis"
	"quants/config"
)

var ctx = context.Background()

func TestNewRepository(t *testing.T) {
	if err := flag.Set("cf", "../../../config/config.yaml"); err != nil {
		panic(err)
	}
	config.Init()
	Init(
		WithMySQL(),
		WithRedis(),
	)
	// mysql
	mysql.NewExample(Clients.MySQL)
	assert.NotNil(t, Example)
	assert.NotNil(t, Example.GetDB(ctx))

	// redis
	redis.NewHealthCheck(Clients.Redis)
	assert.NotNil(t, HealthCheck)
	err := HealthCheck.HealthCheck(ctx)
	assert.Nil(t, err)
}
