package sms

import (
	"fmt"
	"iw/v2/internal"
)

func Start() {

	// TODO: 通过设备获取上下文
	gw := &Sms{}
	var op = internal.BindWallet

	defer func() {
		// 异常统一处理
		recover()
		to := gw.getSender()
		msg := fmt.Sprintf(NetworkErrorAndTryAgainLate, "3 minutes")
		gw.SendMessage(&to, &msg)
	}()

	var cmd internal.Receiver

	switch op {
	case internal.BindWallet:
		cmd = &BindWalletCommand{Terminal: gw}
	case internal.CheckBalance:
		cmd = &CheckBalanceCommand{Terminal: gw}
	case internal.QueryGateway:
		cmd = &QueryGatewayCommand{Terminal: gw}
	case internal.QueryTransaction:
		cmd = &QueryTransactionCommand{Terminal: gw}
	case internal.Transfer:
		cmd = &TransferCommand{Terminal: gw}
	}

	cmd.Execute()
}
