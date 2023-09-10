package sms

import (
	"iw/v2/internal"
	"iw/v2/internal/wallet"
	"strings"
)

type QueryTransactionCommand struct {
	Terminal internal.Terminal
}

func (c *QueryTransactionCommand) Execute() {
	c.Terminal.OnQueryTransaction()
}

func (sms *Sms) OnQueryTransaction() {
	phoneNumber := sms.getSender()

	addr, _, exists := wallet.FindAddress(&phoneNumber)

	if exists {
		if balance, transactions, err := wallet.QueryTransactions(addr); err != nil {
			msg := err.Error()
			sms.SendMessage(&phoneNumber, &msg)
		} else {
			msg := &strings.Builder{}
			msg.WriteString(balance + " ")
			for i := 0; i < 10 && i < len(transactions); i++ {
				msg.WriteString(transactions[i] + " ")
			}
			s := msg.String()
			sms.SendMessage(&phoneNumber, &s)
		}
	} else {
		msg := NoWallet
		sms.SendMessage(&phoneNumber, &msg)
	}
}
