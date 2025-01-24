package repo

import (
	"context"
)

type ITrader interface {
	Buy(ctx context.Context, symbol string, buyPrice, quantity float64) bool
	Sell(ctx context.Context, symbol string, sellPrice, quantity float64) bool
}
