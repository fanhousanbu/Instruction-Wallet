package sms

import (
	"fmt"
	"iw/v2/internal"
	"iw/v2/internal/wallet"
)

// BindWalletCommand 绑定钱包命令
type BindWalletCommand struct {
	Terminal internal.Terminal
	Address  []string // 钱包地址列表
}

// Execute 执行绑定钱包指令
func (c *BindWalletCommand) Execute() {
	c.Terminal.OnBindWallet()
}

func (sms *Sms) OnBindWallet() {
	phoneNumber := sms.getSender()

	msg := SuccessfulAndWaitToConfirm
	sms.SendMessage(&phoneNumber, &msg)

	addr, sign, exists := wallet.FindAddress(&phoneNumber)
	sms.addr = addr
	sms.sign = sign

	if exists {
		msg = fmt.Sprintf(MobileExists, *sms.addr)
		sms.SendMessage(&phoneNumber, &msg)
	} else {
		msg = fmt.Sprintf(MobileBinding, *sms.addr)
		sms.SendMessage(&phoneNumber, sms.addr)
	}
}
