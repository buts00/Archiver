package vlc

import (
	"strings"
	"unicode"
)

type EncoderDecoder struct{}

func New() EncoderDecoder {
	return EncoderDecoder{}
}

// EncodingTable represents the encoding table for converting characters to binary.
type EncodingTable map[rune]string

// Encode encodes the given data string into hexadecimal format.
func (_ EncoderDecoder) Encode(data string) []byte {
	data = PrepareData(data)
	chunks := SplitByChunks(EncodeToBin(data), chunkSize)
	return chunks.Binary()
}

// Decode decodes the given decoded data  back into its original form.
func (_ EncoderDecoder) Decode(data []byte) string {
	binData := NewBinChunks(data).Join()

	decodingTree := getEncodingTable().BuildDecodingTree()

	return ExportData(decodingTree.Decode(binData))
}

// EncodeToBin encodes the data string into binary format.
func EncodeToBin(data string) string {
	builder := strings.Builder{}
	for _, ch := range data {
		builder.WriteString(bin(ch))
	}
	return builder.String()
}

// bin returns the binary representation of the given character.
func bin(ch rune) string {
	table := getEncodingTable()
	val, ok := table[ch]
	if !ok {
		panic("unknown symbol " + string(ch))
	}
	return val
}

// PrepareData prepares the data string for encoding.
func PrepareData(data string) string {
	builder := strings.Builder{}
	for _, ch := range data {
		if unicode.IsUpper(ch) {
			builder.WriteRune('!')
		}
		builder.WriteRune(unicode.ToLower(ch))
	}
	return builder.String()
}

// ExportData exports the decoded data back into its original form by converting the marked uppercase letters to uppercase.
func ExportData(data string) string {
	var (
		builder strings.Builder
		isCaps  bool
	)

	for _, ch := range data {
		if isCaps {
			builder.WriteRune(unicode.ToUpper(ch))
			isCaps = false
			continue
		}
		if ch == '!' {
			isCaps = true
			continue
		}
		builder.WriteRune(ch)
	}

	return builder.String()
}

// getEncodingTable returns the encoding table used for encoding characters into binary.
func getEncodingTable() EncodingTable {
	return EncodingTable{
		' ': "11",
		'e': "101",
		't': "1001",
		'o': "10001",
		'n': "10000",
		'a': "011",
		's': "0101",
		'i': "01001",
		'r': "01000",
		'h': "0011",
		'd': "00101",
		'l': "001001",
		'!': "001000",
		'u': "00011",
		'c': "000101",
		'f': "000100",
		'm': "000011",
		'p': "0000101",
		'g': "0000100",
		'w': "0000011",
		'b': "0000010",
		'y': "0000001",
		'v': "00000001",
		'j': "000000001",
		'k': "0000000001",
		'x': "00000000001",
		'q': "000000000001",
		'z': "000000000000",
	}
}
