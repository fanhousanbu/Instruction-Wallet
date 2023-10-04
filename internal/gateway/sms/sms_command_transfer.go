package sms

import (
	"errors"
	"iw/v2/internal"
	"iw/v2/internal/wallet"
)

type TransferCommand struct {
	Terminal         internal.Terminal
	dstWalletAddress *string // 目标钱包地址
	balance          string  // 转账金额
}

func (c *TransferCommand) Execute() {
	c.Terminal.OnTransfer(c.dstWalletAddress, c.balance)
}

func (sms *Sms) OnTransfer(dst *string, balance string) {
	from := sms.getSender()

	ok, err := wallet.Transfer(&from, dst, balance)

	if ok {
		msg := SuccessfulAndWaitToConfirm
		sms.Reply(&msg)
	} else {
		var e *InsufficientBalanceError
		if errors.Is(err, e) {
			msg := err.Error()
			sms.Reply(&msg)
		} else {
			panic(err) // 抛出异常，由外部统一处理
		}
	}
}
