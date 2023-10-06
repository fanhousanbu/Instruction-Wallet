package instructions

import (
	"iw/v2/internal"
	"iw/v2/internal/instructions/generic"
	"iw/v2/internal/wallet"
	"strings"
)

type QueryTransaction struct {
	generic.BaseProcessor
}

func (c *QueryTransaction) Execute() {
	addr, _, exists := wallet.FindAddressByToken(c.GetToken())

	if exists {
		if balance, transactions, err := wallet.QueryTransactions(addr); err != nil {
			msg := err.Error()
			c.Terminal.Reply(&msg)
		} else {
			msg := &strings.Builder{}
			msg.WriteString(balance + " ")
			for i := 0; i < 10 && i < len(transactions); i++ {
				msg.WriteString(transactions[i] + " ")
			}
			s := msg.String()
			c.Terminal.Reply(&s)
		}
	} else {
		msg := internal.NoWallet
		c.Terminal.Reply(&msg)
	}
}
