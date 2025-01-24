package service

import (
	"context"
	"testing"

	"quants/adapter/repository"
)

var ctx = context.Background()

func TestMain(m *testing.M) {
	repository.Init(
		repository.WithMySQL(),
		repository.WithRedis(),
	)
	m.Run()
}
