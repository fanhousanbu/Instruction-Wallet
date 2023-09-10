package sms

type mockGateway struct {
	from string
}

// GenPhoneNumber 模拟生成手机号
func (m *mockGateway) GenPhoneNumber() string {
	return "13888888888"
}
