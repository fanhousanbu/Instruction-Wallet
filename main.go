package main

import (
	"iw/v2/internal/gateway/sms"
)

func main() {

	// TODO: 启动后端监听服务用来接收终端指令

	sms.Start()
}
