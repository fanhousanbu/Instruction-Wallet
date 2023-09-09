package sms

import "iw/v2/internal"

// BindWalletCommand 绑定钱包命令
type BindWalletCommand struct {
	Terminal internal.Terminal
	Address  []string // 钱包地址列表
}

// Execute 执行绑定钱包指令
func (c *BindWalletCommand) Execute() {
	c.Terminal.OnBindWallet()
}

// Send 发送短信
func (c *BindWalletCommand) Send() {
	c.Terminal.SendMessage(&c.Address[0])
}
