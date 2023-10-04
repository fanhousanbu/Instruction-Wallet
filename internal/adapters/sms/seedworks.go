package sms

const (
	BindWallet          = "0"               // 绑定钱包指令
	Transfer            = `^1 [0-9]{3,15}$` // 转账
	QueryTransaction    = "qt"              // 查询最近10条交易
	SendMyWalletAddress = "g"               // 发送我的钱包地址
)
