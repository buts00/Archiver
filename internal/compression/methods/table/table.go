package table

import "strings"

type EncodingTable map[rune]string

func (et EncodingTable) Decode(data string) string {
	dt := et.buildDecodingTree()
	return dt.Decode(data)
}

type Generator interface {
	NewTable(string2 string) EncodingTable
}

type decodingTree struct {
	Value string
	Zero  *decodingTree
	One   *decodingTree
}

func (dt *decodingTree) add(key rune, value string) {
	currNode := dt

	for _, ch := range value {
		switch ch {
		case '0':
			if currNode.Zero == nil {
				currNode.Zero = &decodingTree{}
			}
			currNode = currNode.Zero
		case '1':
			if currNode.One == nil {
				currNode.One = &decodingTree{}
			}
			currNode = currNode.One
		}
	}
	currNode.Value = string(key)
}

func (dt *decodingTree) Decode(str string) string {
	var builder strings.Builder
	curNode := dt

	for _, ch := range str {
		if curNode.Value != "" {
			builder.WriteString(curNode.Value)

			curNode = dt
		}
		switch ch {
		case '0':
			curNode = curNode.Zero
		case '1':
			curNode = curNode.One
		}
	}

	builder.WriteString(curNode.Value)

	return builder.String()
}

func (et EncodingTable) buildDecodingTree() decodingTree {
	res := decodingTree{}

	for ch, value := range et {
		res.add(ch, value)
	}

	return res
}
