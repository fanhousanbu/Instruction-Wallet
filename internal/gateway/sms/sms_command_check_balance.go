package sms

import "iw/v2/internal"

// CheckBalanceCommand 查询余额命令
type CheckBalanceCommand struct {
	Terminal internal.Terminal
}

func (c *CheckBalanceCommand) Execute() {

}
