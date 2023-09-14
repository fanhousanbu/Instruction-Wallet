package sms

const (
	QueryGateway        = "q"       // 查询可用网关地址
	BindWallet          = "0"       // 绑定钱包指令
	Transfer            = "1"       // 转账
	QueryTransaction    = "qt"      // 查询最近10条交易
	SendMyWalletAddress = "g"       // 发送我的钱包地址
	Unknown             = "unknown" // 未知指令
)
