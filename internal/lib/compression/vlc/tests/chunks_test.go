package tests

import (
	"github.com/buts00/Archiver/internal/lib/compression/vlc"
	"reflect"
	"testing"
)

func Test_splitByChunks(t *testing.T) {
	type args struct {
		data      string
		chunkSize int
	}
	tests := []struct {
		name string
		args args
		want vlc.BinaryChunks
	}{
		{
			name: "base case",
			args: args{
				data:      "001000100110100101",
				chunkSize: 8,
			},
			want: vlc.BinaryChunks{"00100010", "01101001", "01000000"},
		},
		{
			name: "longer string",
			args: args{
				data:      "001000100110100101100011",
				chunkSize: 8,
			},
			want: vlc.BinaryChunks{"00100010", "01101001", "01100011"},
		},
		{
			name: "exact chunk size",
			args: args{
				data:      "00100010",
				chunkSize: 8,
			},
			want: vlc.BinaryChunks{"00100010"},
		},
		{
			name: "empty string",
			args: args{
				data:      "",
				chunkSize: 8,
			},
			want: vlc.BinaryChunks{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := vlc.SplitByChunks(tt.args.data, tt.args.chunkSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitByChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBinChunks(t *testing.T) {

	tests := []struct {
		name string
		data []byte
		want vlc.BinaryChunks
	}{
		{
			name: "base case",
			data: []byte{20, 30, 60, 18},
			want: vlc.BinaryChunks{"00010100", "00011110", "00111100", "00010010"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := vlc.NewBinChunks(tt.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBinChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryChunks_Join(t *testing.T) {
	tests := []struct {
		name string
		bcs  vlc.BinaryChunks
		want string
	}{
		{
			name: "base case",
			bcs:  vlc.BinaryChunks{"00101111", "10000000"},
			want: "0010111110000000",
		},
		{
			name: "addition case",
			bcs:  vlc.BinaryChunks{"00000000", "00100000", "01000000", "00000000"},
			want: "00000000001000000100000000000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bcs.Join(); got != tt.want {
				t.Errorf("Join() = %v, want %v", got, tt.want)
			}
		})
	}
}
