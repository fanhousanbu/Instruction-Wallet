package sms

const (
	ReceiverInstructionMessage  = "successfully and wait to confirm" // 收到绑定钱包的指令，并等待确定
	MobileExists                = "your mobile is already assigned with a wallet address: %s"
	MobileBinding               = "your mobile get a binding wallet address: %s"
	NetworkErrorAndTryAgainLate = "Network error, please try again after %s"
	NotFound                    = "not found any gateways" // 没有找到可用网关
)
