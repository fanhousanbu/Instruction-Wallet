package sms

import (
	"iw/v2/internal/errors"
	"iw/v2/internal/instructions"
	"iw/v2/internal/instructions/generic"
	"regexp"
	"strings"
)

type Adapter struct {
	message *string // 短信内容
	origin  *string // 发信号码
}

// GetProcessor 根据短信内容获取指令处理器
func (sms *Adapter) GetProcessor() (generic.Processor, error) {
	if strings.EqualFold(*sms.message, QueryTransaction) {
		return &instructions.QueryTransaction{
			BaseProcessor: generic.BaseProcessor{RawId: sms.origin, Terminal: sms},
		}, nil
	}

	if strings.EqualFold(*sms.message, BindWallet) {
		return &instructions.BindWalletCommand{
			BaseProcessor: generic.BaseProcessor{RawId: sms.origin, Terminal: sms},
		}, nil
	}

	if strings.EqualFold(*sms.message, SendMyWalletAddress) {
		return &instructions.SendWalletAddress{
			BaseProcessor: generic.BaseProcessor{RawId: sms.origin, Terminal: sms},
		}, nil
	}

	re := regexp.MustCompile(Transfer)
	if match := re.MatchString(*sms.message); match {
		return &instructions.Transfer{
			BaseProcessor: generic.BaseProcessor{RawId: sms.origin, Terminal: sms},
		}, nil
	}

	return nil, errors.UnknownInstructionError{InstructionName: *sms.message}
}
