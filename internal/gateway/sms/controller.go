package sms

import "iw/v2/internal"

func Start() {
	var op = internal.BindWallet

	switch op {
	case internal.BindWallet:
		gw.OnBindWallet()
	case internal.CheckBalance:
		gw.OnCheckBalance()
	case internal.QueryGateway:
		gw.OnQueryGateway()
	case internal.QueryTransaction:
		gw.OnQueryTransaction()
	case internal.Transfer:
		gw.OnTransfer()
	default:
		panic("unknown command")
	}
}
