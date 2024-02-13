package shannon_fano

import (
	"github.com/buts00/Archiver/internal/compression/methods/table"
	"reflect"
	"testing"
)

func Test_bestDividerPosition(t *testing.T) {

	tests := []struct {
		name  string
		codes []code
		want  int
	}{
		{
			name: "one element",
			codes: []code{
				{Quantity: 2},
			},
			want: 0,
		},
		{
			name: "two elements",
			codes: []code{
				{Quantity: 2}, {Quantity: 2},
			},
			want: 1,
		},
		{
			name: "three elements",
			codes: []code{
				{Quantity: 2}, {Quantity: 1}, {Quantity: 1},
			},
			want: 1,
		},
		{
			name: "many elements1",
			codes: []code{
				{Quantity: 2}, {Quantity: 2}, {Quantity: 1}, {Quantity: 1}, {Quantity: 1},
			},
			want: 2,
		},
		{
			name: "uncertainty 3 elements (need leftmost)",
			codes: []code{
				{Quantity: 1}, {Quantity: 1}, {Quantity: 1},
			},
			want: 1,
		},
		{
			name: "uncertainty 4 elements (need leftmost)",
			codes: []code{
				{Quantity: 2}, {Quantity: 2}, {Quantity: 1}, {Quantity: 1},
			},
			want: 1,
		},
		{
			name: "many elements2",
			codes: []code{
				{Quantity: 4},
				{Quantity: 2},
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bestDividerPosition(tt.codes); got != tt.want {
				t.Errorf("bestDividerPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_assignCode(t *testing.T) {

	tests := []struct {
		name  string
		codes []code
		want  []code
	}{
		{
			name:  "empty slice",
			codes: []code{},
			want:  []code{},
		},
		{
			name:  "two elements",
			codes: []code{{Quantity: 2}, {Quantity: 2}},
			want:  []code{{Quantity: 2, Bits: 0, Size: 1}, {Quantity: 2, Bits: 1, Size: 1}},
		},
		{
			name:  "three elements certain pos",
			codes: []code{{Quantity: 2}, {Quantity: 1}, {Quantity: 1}},
			want:  []code{{Quantity: 2, Bits: 0, Size: 1}, {Quantity: 1, Bits: 2, Size: 2}, {Quantity: 1, Bits: 3, Size: 2}},
		},
		{
			name:  "three elements uncertain pos",
			codes: []code{{Quantity: 1}, {Quantity: 1}, {Quantity: 1}},
			want:  []code{{Quantity: 1, Bits: 0, Size: 1}, {Quantity: 1, Bits: 2, Size: 2}, {Quantity: 1, Bits: 3, Size: 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assignCode(tt.codes)

			if !reflect.DeepEqual(tt.codes, tt.want) {
				t.Errorf("got: %v, want: %v", tt.codes, tt.want)
			}
		})
	}
}

func Test_build(t *testing.T) {

	tests := []struct {
		name string
		text string
		want encodingTable
	}{
		{
			name: "base case",
			text: "abbbcc",
			want: encodingTable{
				'a': {
					Char:     'a',
					Quantity: 1,
					Bits:     3,
					Size:     2,
				},
				'b': {
					Char:     'b',
					Quantity: 3,
					Bits:     0,
					Size:     1,
				},
				'c': {
					Char:     'c',
					Quantity: 2,
					Bits:     2,
					Size:     2,
				},
			},
		},
		{
			name: "equal quantity",
			text: "aabbcc",
			want: encodingTable{
				'a': {
					Char:     'a',
					Quantity: 2,
					Bits:     0,
					Size:     1,
				},
				'b': {
					Char:     'b',
					Quantity: 2,
					Bits:     2,
					Size:     2,
				},
				'c': {
					Char:     'c',
					Quantity: 2,
					Bits:     3,
					Size:     2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := build(newCharStat(tt.text)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("build() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerator_NewTable(t *testing.T) {

	tests := []struct {
		name string
		data string
		want table.EncodingTable
	}{
		{
			name: "base case",
			data: "abbbcc",
			want: table.EncodingTable{
				'a': "11",
				'b': "0",
				'c': "10",
			},
		},
		{
			name: "equal quantity",
			data: "aabbcc",
			want: table.EncodingTable{
				'a': "0",
				'b': "10",
				'c': "11",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Generator{}
			if got := g.NewTable(tt.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTable() = %v, want %v", got, tt.want)
			}
		})
	}
}
