package tests

import (
	"github.com/buts00/Archiver/internal/lib/compression/vlc"
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
			if got := vlc.PrepareData(tt.data); got != tt.want {
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
			if got := vlc.EncodeToBin(tt.data); got != tt.want {
				t.Errorf("encodeToBin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncode(t *testing.T) {

	tests := []struct {
		name string
		data string
		want []byte
	}{
		{
			name: "base case",
			data: "My name is Ted",
			want: []byte{32, 48, 60, 24, 119, 74, 228, 77, 40},
		},
	}
	for _, tt := range tests {
		encoder := vlc.New()
		t.Run(tt.name, func(t *testing.T) {
			if got := encoder.Encode(tt.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name        string
		encodedText []byte
		want        string
	}{
		{
			name:        "base case",
			encodedText: []byte{32, 48, 60, 24, 119, 74, 228, 77, 40},
			want:        "My name is Ted",
		},
	}
	for _, tt := range tests {
		decoder := vlc.New()
		t.Run(tt.name, func(t *testing.T) {
			if got := decoder.Decode(tt.encodedText); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
