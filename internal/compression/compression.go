package compression

type Encoder interface {
	Encode(data string) []byte
}

type Decoder interface {
	Decode([]byte) string
}
