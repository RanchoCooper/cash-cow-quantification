package mysql

import (
	"context"
	"testing"

	"quants/domain/entity"
)

var ctx = context.Background()

func TestMain(m *testing.M) {
	db := NewExample(NewMySQLClient()).GetDB(ctx)
	_ = db.AutoMigrate(&entity.Example{})
	m.Run()
}
