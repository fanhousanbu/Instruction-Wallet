package sms

import (
	"iw/v2/internal"
	"iw/v2/internal/wallet"
)

type SendWalletAddressCommand struct {
	Terminal         internal.Terminal
	dstWalletAddress *string // 目标钱包地址
}

func (c *SendWalletAddressCommand) Execute() {
	c.Terminal.OnSendWalletAddressCommand(c.dstWalletAddress)
}

func (sms *Sms) OnSendWalletAddressCommand(dst *string) {
	from := sms.getSender()

	addr, _, exists := wallet.FindAddress(&from)

	if exists {
		sms.SendMessage(dst, addr)
	} else {
		msg := NoWallet
		sms.SendMessage(&from, &msg)
	}
}
