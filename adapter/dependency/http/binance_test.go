package http

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"quants/domain/vo"
)

var ctx = context.Background()

func TestPing(t *testing.T) {
	assert.Equal(t, &vo.PingResp{}, BinanceClinet.Ping(ctx))
}

func TestGetTickerPrice(t *testing.T) {
	result := BinanceClinet.GetTickerPrice(ctx, "ETHBTC")
	fmt.Println(result)
	assert.NotEmpty(t, result)
}

func TestGetTicker24Hour(t *testing.T) {
	result := BinanceClinet.GetTicker24Hour(ctx, "ETHBTC")
	fmt.Println(result)
	assert.NotEmpty(t, result)
}

func TestGetTickerKLine(t *testing.T) {
	result := BinanceClinet.GetTickerKLine(ctx, "ETHBTC", "1M", 20, 1609488000, 1641024000)
	fmt.Println(result)
	assert.NotEmpty(t, result)
}

func TestTradeLimit(t *testing.T) {
	quantity := 1.0
	price := 0.00000001
	result := BinanceClinet.TradeLimit(ctx, "ETHBTC", "BUY", &quantity, &price)
	// {"code":-1021,"msg":"Timestamp for this request is outside of the recvWindow."}
	fmt.Println(result)
	// assert.NotEmpty(t, result)
}
