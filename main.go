package main

import (
	"fmt"
	"github.com/tarm/serial"
	"iw/v2/internal/adapters/sms"
)

func main() {

	// TODO: 启动后端监听服务用来接收终端指令
	if err := sms.Init(&serial.Config{
		Name:        "",
		Baud:        0,
		ReadTimeout: 0,
		Size:        0,
		Parity:      0,
		StopBits:    0,
	}); err != nil {
		fmt.Println("初始化短信网关失败")
	}
}
