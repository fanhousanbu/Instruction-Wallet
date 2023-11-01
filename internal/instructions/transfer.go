package instructions

import (
	"errors"
	"github.com/Microsoft/go-winio/pkg/guid"
	"iw/v2/internal"
	errors2 "iw/v2/internal/errors"
	"iw/v2/internal/instructions/generic"
	"iw/v2/internal/wallet"
)

type Transfer struct {
	generic.BaseProcessor
	id                guid.GUID // 交易ID
	CounterPartyToken *string
	Balance           string
}

func (c *Transfer) Execute() {
	ok, id, err := wallet.TransferByToken(c.GetToken(), c.CounterPartyToken, c.Balance)

	if ok {
		msg := internal.SuccessfulAndWaitToConfirm
		c.Terminal.Reply(&msg)

		// 启用协程在后台定时轮循交易结果
		go WaitingForTransferResult(*id)
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
