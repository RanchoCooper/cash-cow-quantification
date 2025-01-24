package application

import (
	"testing"

	"quants/config"
	service2 "quants/domain/service"
	"quants/domain/strategy/spot_trend_grid"
)

func TestSpotTrendGridLoop(t *testing.T) {
	service2.Init(ctx)
	realTrader := spot_trend_grid.NewTrader()
	simulateTrader := service2.NewSimulatorService(ctx)
	if config.Config.Env == "local" {
		t.Run("simulate", func(t *testing.T) {
			SpotTrendGridLoop(ctx, simulateTrader)
		})

		t.Run("real", func(t *testing.T) {
			SpotTrendGridLoop(ctx, realTrader)
		})
	}
}
