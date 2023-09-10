package sms

import (
	"iw/v2/internal"
	"strings"
)

// QueryGatewayCommand 查询见着命令
type QueryGatewayCommand struct {
	Terminal internal.Terminal
}

func (c *QueryGatewayCommand) Execute() {
	c.Terminal.OnQueryGateway()
}

func (sms *Sms) OnQueryGateway() {
	phoneNumber := sms.getSender()

	if sms.nextGateway >= len(sms.gateways) {
		sms.nextGateway = 0
	}

	var gatewaysList strings.Builder
	for ; sms.nextGateway < len(sms.gateways) && sms.nextGateway < MaxGatewaysList; sms.nextGateway++ {
		gatewaysList.WriteString(sms.gateways[sms.nextGateway] + " ")
	}
	if gatewaysList.Len() > 0 {
		msg := gatewaysList.String()
		sms.SendMessage(&phoneNumber, &msg)
	} else {
		msg := NotFound
		sms.SendMessage(&phoneNumber, &msg)
	}
}
