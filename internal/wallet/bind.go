package wallet

import (
	"iw/v2/internal/proxy"
)

func Bind(token *string) (addr *string, err error) {
	baseUri := "mock.community.node"
	agent := proxy.Agent{BaseApi: &baseUri}
	if addr, exists, err := agent.GetWalletAddressAgent(token); err != nil {
		return nil, err
	} else if exists {
		return addr, nil
	}

	return agent.BindWalletAddressAgent(token)
}
