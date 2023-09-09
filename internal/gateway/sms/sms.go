package sms

import (
	"fmt"
	"iw/v2/internal/wallet"
)

type Sms struct {
	// TODO: 短信网关 对象
	mocker mockGateway
	addr   *string
	sign   *string
}

// getSender 获取当前短信发送号码
func (sms *Sms) getSender() string {
	sms.mocker.sender = sms.mocker.GenPhoneNumber()
	return sms.mocker.sender
}

func (sms *Sms) OnQueryGateway() {

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
