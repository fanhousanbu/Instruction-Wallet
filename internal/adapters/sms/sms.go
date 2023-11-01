package sms

import (
	"github.com/tarm/serial"
	"iw/v2/internal/errors"
	"iw/v2/internal/instructions"
	encoder2 "iw/v2/internal/instructions/encoder"
	"iw/v2/internal/instructions/generic"
	"regexp"
	"strings"
)

// Adapter 短信网关适配器
// 一条短信对应一个适配器，一个适配器负责解析并执行指令，以及根据指令要求回复消息
type Adapter struct {
	message *string       // 短信内容
	origin  *string       // 发信号码
	config  serial.Config // 串口配置
}

// GetProcessor 根据短信内容获取指令处理器
func (sms *Adapter) GetProcessor() (generic.Processor, error) {
	var encoder encoder2.HashEncode = &encoder2.Keccak256Hash{}

	if strings.EqualFold(*sms.message, QueryTransaction) {
		return &instructions.QueryTransaction{
			BaseProcessor: generic.BaseProcessor{RawId: sms.origin, Terminal: sms, TokenEncoder: encoder},
		}, nil
	}

	if strings.EqualFold(*sms.message, BindWallet) {
		return &instructions.BindWalletCommand{
			BaseProcessor: generic.BaseProcessor{RawId: sms.origin, Terminal: sms, TokenEncoder: encoder},
		}, nil
	}

	if strings.EqualFold(*sms.message, SendMyWalletAddress) {
		return &instructions.SendWalletAddress{
			BaseProcessor: generic.BaseProcessor{RawId: sms.origin, Terminal: sms, TokenEncoder: encoder},
		}, nil
	}

	re := regexp.MustCompile(Transfer)
	if match := re.MatchString(*sms.message); match {
		return &instructions.Transfer{
			BaseProcessor: generic.BaseProcessor{RawId: sms.origin, Terminal: sms, TokenEncoder: encoder},
		}, nil
	}

	return nil, errors.UnknownInstructionError{InstructionName: *sms.message}
}
