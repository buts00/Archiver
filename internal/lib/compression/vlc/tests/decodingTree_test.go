package tests

import (
	vlc2 "github.com/buts00/Archiver/internal/lib/compression/vlc"
	"reflect"
	"testing"
)

func Test_encodingTable_BuildDecodingTree(t *testing.T) {
	tests := []struct {
		name string
		et   vlc2.EncodingTable
		want vlc2.DecodingTree
	}{
		{
			name: "base tree case",
			et: vlc2.EncodingTable{
				'a': "11",
				'b': "1001",
				'z': "0101",
			},
			want: vlc2.DecodingTree{
				Zero: &vlc2.DecodingTree{
					One: &vlc2.DecodingTree{
						Zero: &vlc2.DecodingTree{
							One: &vlc2.DecodingTree{
								Value: "z",
							},
						},
					},
				},
				One: &vlc2.DecodingTree{
					One: &vlc2.DecodingTree{
						Value: "a",
					},
					Zero: &vlc2.DecodingTree{
						Zero: &vlc2.DecodingTree{
							One: &vlc2.DecodingTree{
								Value: "b",
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.et.BuildDecodingTree(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildDecodingTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
