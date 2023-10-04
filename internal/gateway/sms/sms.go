package sms

import (
	"regexp"
	"strings"
)

// MaxGatewaysList 最多一次获取的网关数量
const MaxGatewaysList int = 3

type Sms struct {
	// TODO: 短信网关 对象
	mocker      mockGateway
	rawMessage  *string  // 原始收取信息
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

func (sms *Sms) GetInstruction() string {
	if strings.EqualFold(*sms.rawMessage, QueryTransaction) {
		return QueryTransaction
	}

	if strings.EqualFold(*sms.rawMessage, QueryGateway) {
		return QueryGateway
	}

	if strings.EqualFold(*sms.rawMessage, BindWallet) {
		return BindWallet
	}

	if strings.EqualFold(*sms.rawMessage, SendMyWalletAddress) {
		return SendMyWalletAddress
	}

	re := regexp.MustCompile(`^1 [0-9]{3,15}$`)
	if match := re.MatchString(*sms.rawMessage); match {
		return Transfer
	}

	return Unknown
}

func (sms *Sms) Reply(msg *string) {
}
