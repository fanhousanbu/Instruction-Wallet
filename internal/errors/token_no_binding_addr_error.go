package errors

// TokenNoBindingAddrError 没有绑定钱包
type TokenNoBindingAddrError struct {
	InstructionName string
}

func (e TokenNoBindingAddrError) Error() string {
	return "token not binding wallet address yet"
}
