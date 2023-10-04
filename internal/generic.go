package internal

// Receiver 接收指令
type Receiver interface {
	Execute()
}

// Terminal 终端对象
// 比如智能手机、PC、传统非智能手机等
type Terminal interface {
	Reply(*string)                          // 回复消息
	OnQueryGateway()                        // 查询网关
	OnBindWallet()                          // 绑定钱包
	OnTransfer(dst *string, balance string) // 转账, dst: 目标地址，balance：转账金额
	OnQueryTransaction()                    // 查询钱包余额及最近10条交易
	OnSendWalletAddress(dst *string)        // 发送我的钱包地址
}
