package redis

import (
	"context"
	"errors"

	"quants/domain/repo"
)

type HealthCheck struct {
	IRedis
}

var _ repo.IHealthCheckRepository = &HealthCheck{}

func NewHealthCheck(redis IRedis) *HealthCheck {
	return &HealthCheck{IRedis: redis}
}

func (h HealthCheck) HealthCheck(ctx context.Context) error {
	pong := h.GetClient().Ping(ctx).String()
	if pong != "ping: PONG" {
		return errors.New("ping redis got invalid response: " + pong)
	}
	return nil
}
