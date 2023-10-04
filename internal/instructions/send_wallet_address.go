package instructions

import (
	"iw/v2/internal"
	"iw/v2/internal/instructions/generic"
	"iw/v2/internal/wallet"
)

type SendWalletAddress struct {
	generic.BaseProcessor
}

func (c *SendWalletAddress) Execute() {
	addr, _, exists := wallet.FindAddressByToken(c.Token)

	if exists {
		c.Terminal.Reply(addr)
	} else {
		msg := internal.NoWallet
		c.Terminal.Reply(&msg)
	}
}
