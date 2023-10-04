package main

import (
	"iw/v2/internal/adapters/sms"
)

func main() {

	// TODO: 启动后端监听服务用来接收终端指令

	msg := "0"
	sms.Start(&msg)
}
