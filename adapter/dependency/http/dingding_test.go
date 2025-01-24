package http

import (
	"testing"
)

func TestSendDingDingMessage(t *testing.T) {
	DingDingClient.SendDingDingMessage("testing message", false)
}
