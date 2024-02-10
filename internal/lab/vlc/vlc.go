package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// BinaryChunks represents an array of binary chunks.
type BinaryChunks []BinaryChunk

// ToHex converts BinaryChunks to HexChunks.
func (bcs BinaryChunks) ToHex() HexChunks {
	res := make(HexChunks, 0, len(bcs))
	for _, chunk := range bcs {
		res = append(res, chunk.ToHex())
	}
	return res
}

// BinaryChunk represents a single binary chunk.
type BinaryChunk string

// ToHex converts a BinaryChunk to a HexChunk.
func (bs BinaryChunk) ToHex() HexChunk {
	num, err := strconv.ParseUint(string(bs), 2, chunkSize)
	if err != nil {
		panic("cannot parse to int" + err.Error())
	}
	res := strings.ToUpper(fmt.Sprintf("%x", num))
	if len(res) == 1 {
		res = "0" + res
	}
	return HexChunk(res)
}

// encodingTable represents the encoding table for converting characters to binary.
type encodingTable map[rune]string

// HexChunk represents a single hexadecimal chunk.
type HexChunk string

// HexChunks represents an array of hexadecimal chunks.
type HexChunks []HexChunk

// ToString converts HexChunks to a single string.
func (hcs HexChunks) ToString() string {
	const sep = " "
	if len(hcs) == 0 {
		return ""
	}
	builder := strings.Builder{}
	builder.WriteString(string(hcs[0]))
	for _, chunk := range hcs[1:] {
		builder.WriteString(sep)
		builder.WriteString(string(chunk))
	}
	return builder.String()
}

// chunkSize represents the size of each binary chunk.
const chunkSize = 8

// Encode encodes the given data string into hexadecimal format.
func Encode(data string) string {
	data = prepareData(data)
	chunks := splitByChunks(encodeToBin(data), chunkSize)
	return chunks.ToHex().ToString()
}

// splitByChunks splits the data string into binary chunks of the given size.
func splitByChunks(data string, chunkSize int) BinaryChunks {
	dataLength := utf8.RuneCountInString(data)
	numberOfChunks := dataLength / chunkSize
	if dataLength%chunkSize != 0 {
		numberOfChunks++
	}
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

// encodeToBin encodes the data string into binary format.
func encodeToBin(data string) string {
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
		panic("unknown symbol" + string(ch))
	}
	return val
}

// prepareData prepares the data string for encoding.
func prepareData(data string) string {
	builder := strings.Builder{}
	for _, ch := range data {
		if unicode.IsUpper(ch) {
			builder.WriteRune('!')
		}
		builder.WriteRune(unicode.ToLower(ch))
	}
	return builder.String()
}

// getEncodingTable returns the encoding table used for encoding characters into binary.
func getEncodingTable() encodingTable {
	return encodingTable{
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
