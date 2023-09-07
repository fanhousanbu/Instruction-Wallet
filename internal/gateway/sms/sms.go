package sms

type Sms struct {
	// TOOD: 短信网关 对象
	mocker mockGateway
}

// getSender 获取当前短信发送号码
func (sms *Sms) getSender() string {
	sms.mocker.sender = sms.mocker.GenPhoneNumber()
	return sms.mocker.sender
}

func (sms *Sms) OnQueryGateway() {

}

func (sms *Sms) OnBindWallet() {

}
func (sms *Sms) OnTransfer() {

}

func (sms *Sms) OnCheckBalance() {

}
func (sms *Sms) OnQueryTransaction() {

}
