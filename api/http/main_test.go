package http

import (
	"context"
	"flag"
	"testing"

	"quants/adapter/repository"
	"quants/config"
	"quants/domain/entity"
	"quants/domain/service"
)

var ctx = context.Background()

func TestMain(m *testing.M) {
	if err := flag.Set("cf", "../../config/config.yaml"); err != nil {
		panic(err)
	}
	config.Init()
	repository.Init(repository.WithMySQL(), repository.WithRedis())
	db := repository.Clients.MySQL.GetDB(ctx)
	_ = db.AutoMigrate(&entity.Example{})

	service.Init(ctx)

	m.Run()
}
