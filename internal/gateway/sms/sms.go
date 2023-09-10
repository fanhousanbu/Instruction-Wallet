package sms

// MaxGatewaysList 最多一次获取的网关数量
const MaxGatewaysList int = 3

type Sms struct {
	// TODO: 短信网关 对象
	mocker      mockGateway
	nextGateway int      // 网关索引
	gateways    []string // 所有可用网关
	addr        *string
	sign        *string
}

// getSender 获取当前短信发送号码
func (sms *Sms) getSender() string {
	sms.mocker.from = sms.mocker.GenPhoneNumber()
	return sms.mocker.from
}

func (sms *Sms) OnCheckBalance() {

}
func (sms *Sms) OnQueryTransaction() {

}

func (sms *Sms) SendMessage(to, msg *string) {

}
