package encoder

import (
	"github.com/ethereum/go-ethereum/crypto"
)

type Keccak256Hash struct {
}

func (p *Keccak256Hash) Encode(raw []byte) (*string, error) {
	hash := crypto.Keccak256Hash(raw)

	str := hash.Hex()
	return &str, nil
}
