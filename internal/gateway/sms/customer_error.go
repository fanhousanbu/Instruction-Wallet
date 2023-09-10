package sms

// InsufficientBalanceError 余额不足
type InsufficientBalanceError struct {
	name string
}

func (e *InsufficientBalanceError) Error() string {
	return "Insufficient balance"
}
