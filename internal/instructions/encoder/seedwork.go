package encoder

type HashEncode interface {
	Encode([]byte) (*string, error)
}
