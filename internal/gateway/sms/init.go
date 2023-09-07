package sms

var (
	gw               *Sms
	bindWallet       *BindWalletCommand
	checkBalance     *CheckBalanceCommand
	queryGateway     *QueryGatewayCommand
	transfer         *TransferCommand
	queryTransaction *QueryTransactionCommand
)

func init() {

	gw = &Sms{}

	bindWallet = &BindWalletCommand{Terminal: gw}
	checkBalance = &CheckBalanceCommand{Terminal: gw}
	queryGateway = &QueryGatewayCommand{Terminal: gw}
	transfer = &TransferCommand{Terminal: gw}
	queryTransaction = &QueryTransactionCommand{Terminal: gw}
}
