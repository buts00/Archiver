package methods

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type BinaryChunks []BinaryChunk

type BinaryChunk string

const (
	chunkSize = 8
)

func (bcs BinaryChunks) Join() string {
	var builder strings.Builder
	for _, chunk := range bcs {
		builder.WriteString(string(chunk))
	}
	return builder.String()
}

func (bcs BinaryChunks) Binary() []byte {
	res := make([]byte, 0, len(bcs))

	for _, chunk := range bcs {
		res = append(res, chunk.Binary())
	}

	return res
}

func (bc BinaryChunk) Binary() byte {
	num, err := strconv.ParseUint(string(bc), 2, chunkSize)
	if err != nil {
		panic("cannot parse to int" + err.Error())
	}
	return byte(num)
}

func NewBinChunks(data []byte) BinaryChunks {
	res := make(BinaryChunks, 0, len(data))

	for _, ch := range data {
		res = append(res, NewBinChunk(ch))
	}

	return res
}

func NewBinChunk(code byte) BinaryChunk {
	return BinaryChunk(fmt.Sprintf("%08b", code))
}

func SplitByChunks(data string, chunkSize int) BinaryChunks {
	dataLength := utf8.RuneCountInString(data)
	numberOfChunks := (dataLength + chunkSize) / chunkSize

	res := make(BinaryChunks, 0, numberOfChunks)
	builder := strings.Builder{}
	for i, ch := range data {
		builder.WriteRune(ch)
		if (i+1)%chunkSize == 0 {
			res = append(res, BinaryChunk(builder.String()))
			builder.Reset()
		}
	}
	if builder.Len() != 0 {
		lastChunk := builder.String()
		lastChunk += strings.Repeat("0", chunkSize-len(lastChunk))
		res = append(res, BinaryChunk(lastChunk))
	}
	return res
}
