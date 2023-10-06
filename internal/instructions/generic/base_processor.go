package generic

import (
	"iw/v2/internal/adapters"
	"iw/v2/internal/instructions/encoder"
)

type BaseProcessor struct {
	RawId        *string
	token        *string
	Terminal     adapters.Terminal
	TokenEncoder encoder.HashEncode
}

func (p *BaseProcessor) GetToken() *string {
	if p.token == nil {
		if t, err := p.TokenEncoder.Encode([]byte(*p.RawId)); err == nil {
			p.token = t
		}
	}
	return p.token
}
