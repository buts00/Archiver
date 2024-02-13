package shannon_fano

import (
	"fmt"
	"github.com/buts00/Archiver/internal/compression/methods/table"
	"math"
	"sort"
	"strings"
)

type Generator struct{}

type charStat map[rune]int

func NewGenerator() Generator {
	return Generator{}
}

type encodingTable map[rune]code

func (et encodingTable) Export() table.EncodingTable {
	res := make(table.EncodingTable)

	for k, code := range et {
		byteStr := fmt.Sprintf("%b", code.Bits)
		byteStr = strings.Repeat("0", code.Size-len(byteStr)) + byteStr
		res[k] = byteStr
	}

	return res
}

type code struct {
	Char     rune
	Quantity int
	Bits     uint32
	Size     int
}

func (g Generator) NewTable(data string) table.EncodingTable {
	return build(newCharStat(data)).Export()
}

func build(stat charStat) encodingTable {
	codes := make([]code, 0, len(stat))

	for ch, qty := range stat {
		codes = append(codes, code{
			Char:     ch,
			Quantity: qty,
		})
	}

	sort.Slice(codes, func(i, j int) bool {
		if codes[i].Quantity != codes[j].Quantity {
			return codes[i].Quantity > codes[j].Quantity
		}
		return codes[i].Char < codes[j].Char
	})

	assignCode(codes)

	res := make(encodingTable)

	for _, code := range codes {
		res[code.Char] = code
	}

	return res
}

func assignCode(codes []code) {
	size := len(codes)
	if size < 2 {
		return
	}

	bestPosition := bestDividerPosition(codes)

	for i := 0; i < size; i++ {
		codes[i].Bits <<= 1
		codes[i].Size++
		if i >= bestPosition {
			codes[i].Bits++
		}
	}

	assignCode(codes[:bestPosition])
	assignCode(codes[bestPosition:])

}

func bestDividerPosition(codes []code) int {
	sum, left, prevDiff, bestPos := 0, 0, math.MaxInt, 0

	for _, ch := range codes {
		sum += ch.Quantity
	}

	for i := 0; i < len(codes)-1; i++ {
		left += codes[i].Quantity
		sum -= codes[i].Quantity
		diff := abs(sum - left)
		if diff >= prevDiff {
			break
		}
		prevDiff = diff
		bestPos = i + 1
	}

	return bestPos
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func newCharStat(data string) charStat {
	res := make(charStat)
	for _, ch := range data {
		res[ch]++
	}
	return res
}
