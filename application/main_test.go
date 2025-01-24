package application

import (
	"context"
	"testing"

	"quants/adapter/repository"
	"quants/domain/service"
)

var ctx = context.Background()

func TestMain(m *testing.M) {
	repository.Init(
		repository.WithMySQL(),
		repository.WithRedis(),
	)
	service.Init(ctx)
	m.Run()
}
