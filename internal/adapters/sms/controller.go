package sms

import (
	"fmt"
	"iw/v2/internal"
)

func Start(msg *string) {

	gw := &Adapter{
		message: msg,
	}

	defer func() {
		// 异常统一处理
		recover()
		msg := fmt.Sprintf(internal.NetworkErrorAndTryAgainLate, "3 minutes")
		gw.Reply(&msg)
	}()

	if op, err := gw.GetProcessor(); err != nil {
		panic(err)
	} else {
		op.Execute()
	}
}
