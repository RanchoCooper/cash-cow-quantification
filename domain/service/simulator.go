package service

import (
	"context"

	"github.com/shopspring/decimal"

	"quants/adapter/repository"
	entity2 "quants/domain/entity"
	"quants/domain/repo"
	"quants/util"
	"quants/util/logger"
)

const (
	SimulatorUserName  = "simulator"
	SimulatorUserEmail = "rancho@simulator.com"
)

type SimulatorService struct {
	UserRepository  repo.IUserRepo
	TradeRepository repo.ITradeRepo
}

func NewSimulatorService(ctx context.Context) *SimulatorService {
	srv := &SimulatorService{UserRepository: repository.User, TradeRepository: repository.Trade}
	logger.Log.Info(ctx, "simulator service init successfully")
	return srv
}

func (ss *SimulatorService) FindOrCreateSimulateUser(ctx context.Context) *entity2.User {
	user, err := ss.UserRepository.FindByEmail(ctx, SimulatorUserEmail)
	if err != nil {
		logger.Log.Errorf(ctx, "find simulate user fail when simulate, err: %v", err)
		return nil
	}
	if user == nil {
		user, err = ss.UserRepository.Create(ctx, nil, &entity2.User{
			UserName:  SimulatorUserName,
			UserEmail: SimulatorUserEmail,
			Asset:     decimal.NewFromFloat(10000),
			Profit:    decimal.NewFromFloat(0),
			State:     entity2.UserStateEnable,
		})
		if err != nil {
			logger.Log.Errorf(ctx, "create simulate user fail when simulate, err: %v", err)
			return nil
		}
	}
	return user
}

func (ss *SimulatorService) Buy(ctx context.Context, symbol string, price, quantity float64) bool {
	user := ss.FindOrCreateSimulateUser(ctx)
	if user == nil {
		return false
	}

	trade := &entity2.Trade{
		OrderId:    util.RandString(20, false),
		UserEmail:  user.UserEmail,
		Symbol:     symbol,
		Type:       entity2.TradeTypeBuy,
		Price:      decimal.NewFromFloat(price),
		Quantity:   decimal.NewFromFloat(quantity),
		IsSimulate: true,
	}
	trade, err := ss.TradeRepository.Create(ctx, nil, trade)
	if err != nil {
		logger.Log.Errorf(ctx, "create user trade fail when simulate, trade: %v, err: %v", trade, err)
		return false
	}

	// update user asset
	user.Asset = user.Asset.Sub(decimal.NewFromFloat(price * quantity))
	err = ss.UserRepository.Update(ctx, nil, user)
	if err != nil {
		logger.Log.Errorf(ctx, "update user fail when simulate, err: %v", err)
		return false
	}
	return true
}

func (ss *SimulatorService) Sell(ctx context.Context, symbol string, price, quantity float64) bool {
	user := ss.FindOrCreateSimulateUser(ctx)
	if user == nil {
		return false
	}

	trade := &entity2.Trade{
		OrderId:    util.RandString(20, false),
		UserEmail:  user.UserEmail,
		Symbol:     symbol,
		Type:       entity2.TradeTypeSell,
		Price:      decimal.NewFromFloat(price),
		Quantity:   decimal.NewFromFloat(quantity),
		IsSimulate: true,
	}
	trade, err := ss.TradeRepository.Create(ctx, nil, trade)
	if err != nil {
		logger.Log.Errorf(ctx, "create user trade fail when simulate, trade: %v, err: %v", trade, err)
		return false
	}

	// update user asset
	user.Asset = user.Asset.Add(decimal.NewFromFloat(price * quantity))
	err = ss.UserRepository.Update(ctx, nil, user)
	if err != nil {
		logger.Log.Errorf(ctx, "update user fail when simulate, err: %v", err)
		return false
	}
	return true
}
