package generic

import "iw/v2/internal/adapters"

type BaseProcessor struct {
	RawId    *string
	Token    *string
	Terminal adapters.Terminal
}

func (p *BaseProcessor) GetToken() *string {
	if p.Token == nil {
		// TODO: poseidon hash
		*p.Token = "*" + *p.RawId
	}
	return p.Token
}
