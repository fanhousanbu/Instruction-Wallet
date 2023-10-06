package encoder

import (
	"github.com/iden3/go-iden3-crypto/poseidon"
)

type PoseidonHash struct {
}

func (p *PoseidonHash) Encode(raw []byte) (*string, error) {
	if h, err := poseidon.HashBytes(raw); err != nil {
		return nil, err
	} else {
		s := h.String()
		return &s, nil
	}
}
