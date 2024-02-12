package vlc

import "strings"

type DecodingTree struct {
	Value string
	Zero  *DecodingTree
	One   *DecodingTree
}

// Add inserts a key-value pair into the decoding tree.
func (dt *DecodingTree) Add(key rune, value string) {
	currNode := dt

	for _, ch := range value {
		switch ch {
		case '0':
			if currNode.Zero == nil {
				currNode.Zero = &DecodingTree{}
			}
			currNode = currNode.Zero
		case '1':
			if currNode.One == nil {
				currNode.One = &DecodingTree{}
			}
			currNode = currNode.One
		}
	}
	currNode.Value = string(key)
}

// Decode decodes the given binary string using the decoding tree and returns the decoded string.
func (dt *DecodingTree) Decode(str string) string {
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

// BuildDecodingTree builds a decoding tree from the encoding table.
func (et EncodingTable) BuildDecodingTree() DecodingTree {
	res := DecodingTree{}

	for ch, value := range et {
		res.Add(ch, value)
	}

	return res
}
