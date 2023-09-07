package internal

const (
	QueryGateway     = "qg" // 查询可用网关地址
	BindWallet       = "0"  // 绑定钱包指令
	Transfer         = "1"  // 转账
	CheckBalance     = "cb" // 查询余额
	QueryTransaction = "qt" // 查询最近10条交易
)

type Command interface {
	Execute()
}

// Terminal represent 终端对象
// 比如智能手机、PC、传统非智能手机等
type Terminal interface {
	OnQueryGateway()     // 查询网关
	OnBindWallet()       // 绑定钱包
	OnTransfer()         // 转账
	OnCheckBalance()     // 查询余额
	OnQueryTransaction() // 查询交易
}
