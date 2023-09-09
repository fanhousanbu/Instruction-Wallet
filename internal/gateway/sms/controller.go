package sms

import "iw/v2/internal"

func Start() {
	var op = internal.BindWallet

	var cmd internal.Receiver
	switch op {
	case internal.BindWallet:
		cmd = bindWallet
	case internal.CheckBalance:
		cmd = checkBalance
	case internal.QueryGateway:
		cmd = queryGateway
	case internal.QueryTransaction:
		cmd = queryTransaction
	case internal.Transfer:
		cmd = transfer
	default:
		panic("unknown command")
	}

	cmd.Execute()
}
