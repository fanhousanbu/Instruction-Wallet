package sms

import (
	"fmt"
	"iw/v2/internal/wallet"
	"strings"
)

// MaxGatewaysList 最多一次获取的网关数量
const MaxGatewaysList int = 3

type Sms struct {
	// TODO: 短信网关 对象
	mocker      mockGateway
	nextGateway int      // 网关索引
	gateways    []string // 所有可用网关
	addr        *string
	sign        *string
}

// getSender 获取当前短信发送号码
func (sms *Sms) getSender() string {
	sms.mocker.from = sms.mocker.GenPhoneNumber()
	return sms.mocker.from
}

func (sms *Sms) OnQueryGateway() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered in QueryGateway", r)
			msg := fmt.Sprintf(NetworkErrorAndTryAgainLate, "3 minutes")
			sms.SendMessage(&msg)
		}
	}()

	if sms.nextGateway >= len(sms.gateways) {
		sms.nextGateway = 0
	}

	var gatewaysList strings.Builder
	for ; sms.nextGateway < len(sms.gateways) || sms.nextGateway < MaxGatewaysList; sms.nextGateway++ {
		gatewaysList.WriteString(sms.gateways[sms.nextGateway] + " ")
	}
	if gatewaysList.Len() > 0 {
		msg := gatewaysList.String()
		sms.SendMessage(&msg)
	} else {
		msg := NotFound
		sms.SendMessage(&msg)
	}
}

func (sms *Sms) OnBindWallet() {
	msg := ReceiverInstructionMessage
	sms.SendMessage(&msg)

	phoneNumber := sms.getSender()
	addr, sign, exists := wallet.FindAddress(&phoneNumber)
	sms.addr = addr
	sms.sign = sign

	if exists {
		msg = fmt.Sprintf(MobileExists, *sms.addr)
		sms.SendMessage(&msg)
	} else {
		msg = fmt.Sprintf(MobileBinding, *sms.addr)
		sms.SendMessage(sms.addr)
	}
}

func (sms *Sms) OnTransfer() {

}

func (sms *Sms) OnCheckBalance() {

}
func (sms *Sms) OnQueryTransaction() {

}

func (sms *Sms) SendMessage(msg *string) {

}
