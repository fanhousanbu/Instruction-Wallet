package wallet

import "github.com/Microsoft/go-winio/pkg/guid"

// TransferByToken 钱包转账
// 返回：bool 是否成功；guid: 如果成功，则表示转账ID
func TransferByToken(from, to *string, balance string) (bool, *guid.GUID, error) {
	return true, nil, nil
}
