package sms

import "iw/v2/internal"

// QueryGatewayCommand 查询见着命令
type QueryGatewayCommand struct {
	Terminal internal.Terminal
}

func (c *QueryGatewayCommand) Execute() {
	c.Terminal.OnQueryGateway()
}

// Send 发送短信
func (c *QueryGatewayCommand) Send() {
}
