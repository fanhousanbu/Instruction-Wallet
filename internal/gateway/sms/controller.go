package sms

import (
	"fmt"
	"iw/v2/internal"
)

func Start() {

	// TODO: 通过设备获取上下文
	msg := "0"
	gw := &Sms{
		rawMessage: &msg,
	}
	var op = gw.GetInstruction()

	defer func() {
		// 异常统一处理
		recover()
		to := gw.getSender()
		msg := fmt.Sprintf(NetworkErrorAndTryAgainLate, "3 minutes")
		gw.SendMessage(&to, &msg)
	}()

	var cmd internal.Receiver

	switch op {
	case BindWallet:
		cmd = &BindWalletCommand{Terminal: gw}
	case QueryGateway:
		cmd = &QueryGatewayCommand{Terminal: gw}
	case QueryTransaction:
		cmd = &QueryTransactionCommand{Terminal: gw}
	case Transfer:
		cmd = &TransferCommand{Terminal: gw}
	case SendMyWalletAddress:
		cmd = &TransferCommand{Terminal: gw}
	}

	cmd.Execute()
}
