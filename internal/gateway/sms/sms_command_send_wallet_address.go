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
	c.Terminal.OnSendWalletAddress(c.dstWalletAddress)
}

func (sms *Sms) OnSendWalletAddress(dst *string) {
	from := sms.getSender()

	addr, _, exists := wallet.FindAddress(&from)

	if exists {
		sms.Reply(addr)
	} else {
		msg := NoWallet
		sms.Reply(&msg)
	}
}
