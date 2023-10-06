package instructions

import (
	"errors"
	"iw/v2/internal"
	errors2 "iw/v2/internal/errors"
	"iw/v2/internal/instructions/generic"
	"iw/v2/internal/wallet"
)

type Transfer struct {
	generic.BaseProcessor
	CounterPartyToken *string
	Balance           string
}

func (c *Transfer) Execute() {
	ok, err := wallet.TransferByToken(c.GetToken(), c.CounterPartyToken, c.Balance)

	if ok {
		msg := internal.SuccessfulAndWaitToConfirm
		c.Terminal.Reply(&msg)
	} else {
		var e *errors2.InsufficientBalanceError
		if errors.Is(err, e) {
			msg := err.Error()
			c.Terminal.Reply(&msg)
		} else {
			panic(err) // 抛出异常，由外部统一处理
		}
	}
}
