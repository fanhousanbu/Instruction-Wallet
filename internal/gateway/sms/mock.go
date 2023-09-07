package sms

type mockGateway struct {
	sender string
}

func (m *mockGateway) GenPhoneNumber() string {
	return "13888888888"
}
