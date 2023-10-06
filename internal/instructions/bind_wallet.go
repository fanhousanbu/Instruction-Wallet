package instructions

import (
	"fmt"
	"iw/v2/internal"
	"iw/v2/internal/instructions/generic"
	"iw/v2/internal/wallet"
)

// BindWalletCommand 绑定钱包命令
type BindWalletCommand struct {
	generic.BaseProcessor
}

// Execute 执行绑定钱包指令
func (c *BindWalletCommand) Execute() {

	msg := internal.SuccessfulAndWaitToConfirm
	c.Terminal.Reply(&msg)

	addr, _, exists := wallet.FindAddressByToken(c.GetToken())

	if exists {
		msg = fmt.Sprintf(internal.MobileExists, *addr)
		c.Terminal.Reply(&msg)
	} else {
		msg = fmt.Sprintf(internal.MobileBinding, *addr)
		c.Terminal.Reply(addr)
	}
}
