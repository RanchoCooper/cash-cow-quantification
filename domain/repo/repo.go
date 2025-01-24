package repo

import (
	"context"

	"gorm.io/gorm"

	entity2 "quants/domain/entity"
)

type IExampleRepo interface {
	Create(ctx context.Context, tx *gorm.DB, entity *entity2.Example) (*entity2.Example, error)
	Delete(ctx context.Context, tx *gorm.DB, Id int) error
	Update(ctx context.Context, tx *gorm.DB, entity *entity2.Example) error
	Get(ctx context.Context, Id int) (entity *entity2.Example, e error)
	FindByName(ctx context.Context, name string) (*entity2.Example, error)
}

type IHealthCheckRepository interface {
	HealthCheck(ctx context.Context) error
}

type ITradeRepo interface {
	Create(ctx context.Context, tx *gorm.DB, entity *entity2.Trade) (*entity2.Trade, error)
	Delete(ctx context.Context, tx *gorm.DB, Id int) error
	Update(ctx context.Context, tx *gorm.DB, entity *entity2.Trade) error
	Get(ctx context.Context, Id int) (entity *entity2.Trade, e error)
	FindByOrderID(context.Context, string) (*entity2.Trade, error)
}

type IUserRepo interface {
	Create(ctx context.Context, tx *gorm.DB, entity *entity2.User) (*entity2.User, error)
	Delete(ctx context.Context, tx *gorm.DB, Id int) error
	Update(ctx context.Context, tx *gorm.DB, entity *entity2.User) error
	Get(ctx context.Context, Id int) (entity *entity2.User, e error)
	FindByEmail(context.Context, string) (*entity2.User, error)
}
