package vlc

import (
	"reflect"
	"testing"
)

func Test_prepareData(t *testing.T) {
	tests := []struct {
		name string
		data string
		want string
	}{
		{
			name: "base case",
			data: "My name is Ted",
			want: "!my name is !ted",
		},
		{
			name: "caps test",
			data: "MY NAME IS TED",
			want: "!m!y !n!a!m!e !i!s !t!e!d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepareData(tt.data); got != tt.want {
				t.Errorf("prepareData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encodeToBin(t *testing.T) {
	tests := []struct {
		name string
		data string
		want string
	}{
		{
			name: "base case",
			data: "!ted",
			want: "001000100110100101",
		},
		{
			name: "empty string",
			data: "",
			want: "",
		},
		{
			name: "single character",
			data: "e",
			want: "101",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeToBin(tt.data); got != tt.want {
				t.Errorf("encodeToBin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitByChunks(t *testing.T) {
	type args struct {
		data      string
		chunkSize int
	}
	tests := []struct {
		name string
		args args
		want BinaryChunks
	}{
		{
			name: "base case",
			args: args{
				data:      "001000100110100101",
				chunkSize: 8,
			},
			want: BinaryChunks{"00100010", "01101001", "01000000"},
		},
		{
			name: "longer string",
			args: args{
				data:      "001000100110100101100011",
				chunkSize: 8,
			},
			want: BinaryChunks{"00100010", "01101001", "01100011"},
		},
		{
			name: "exact chunk size",
			args: args{
				data:      "00100010",
				chunkSize: 8,
			},
			want: BinaryChunks{"00100010"},
		},
		{
			name: "empty string",
			args: args{
				data:      "",
				chunkSize: 8,
			},
			want: BinaryChunks{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitByChunks(tt.args.data, tt.args.chunkSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitByChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryChunks_ToHex(t *testing.T) {
	tests := []struct {
		name string
		bcs  BinaryChunks
		want HexChunks
	}{
		{
			name: "base case",
			bcs:  BinaryChunks{"0101111", "10000000"},
			want: HexChunks{"2F", "80"},
		},
		{
			name: "empty binary chunks",
			bcs:  BinaryChunks{},
			want: HexChunks{},
		},
		{
			name: "binary chunks with leading zeros",
			bcs:  BinaryChunks{"00001111", "00000001"},
			want: HexChunks{"0F", "01"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bcs.ToHex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncode(t *testing.T) {

	tests := []struct {
		name string
		data string
		want string
	}{
		{
			name: "base case",
			data: "My name is Ted",
			want: "20 30 3C 18 77 4A E4 4D 28",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.data); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
