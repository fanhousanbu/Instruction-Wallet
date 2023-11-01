package sms

import (
	"fmt"
	"github.com/tarm/serial"
	"iw/v2/internal"
	"iw/v2/internal/infra"
)

// received 收到短信
func received(msg []byte) {

	m := string(msg)
	gw := &Adapter{
		message: &m,
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

func Init(config *serial.Config) error {
	conn := infra.SmsGatewayConnector{}

	return conn.Listen(config.Name, config.Baud, received)
}
