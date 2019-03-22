package app

import (
	"reflect"
	"testing"

	abciTypes "github.com/tendermint/tendermint/abci/types"
)

var testData = []struct {
	compvals [][32]byte
	pubkeys  [][32]byte
}{
	{
		compvals: [][32]byte{
			[32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x87, 0x86, 0x78, 0x32, 0x6e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		},
		pubkeys: [][32]byte{
			[32]byte{0xf1, 0xa7, 0x75, 0x90, 0xcd, 0xe9, 0x85, 0x99, 0xf1, 0xda, 0x45, 0x7d, 0xf0, 0x10, 0x86, 0x5d, 0xcc, 0x07, 0xa4, 0x6c, 0x4b, 0x26, 0xaa, 0xec, 0xa8, 0x66, 0x8f, 0xf7, 0x29, 0xf7, 0xa5, 0xf9},
		},
	},
	{
		compvals: [][32]byte{
			[32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xa2, 0xa1, 0x5d, 0x09, 0x51, 0xca, 0x35, 0xb7, 0xd9, 0x15, 0x45, 0x8e, 0xf5, 0x40, 0xad, 0xe6, 0x06, 0x8d, 0xfe, 0x2f, 0x44, 0xe8, 0xfa, 0x73, 0x3c},
			[32]byte{0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x87, 0x86, 0x78, 0x32, 0x6e, 0xca, 0x35, 0xb7, 0xd9, 0x15, 0x45, 0x8e, 0xf5, 0x40, 0xad, 0xe6, 0x06, 0x8d, 0xfe, 0x2f, 0x44, 0xe8, 0xfa, 0x73, 0x3c},
			[32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x45, 0x42, 0xba, 0x12, 0xa3, 0xca, 0x35, 0xb7, 0xd9, 0x15, 0x45, 0x8e, 0xf5, 0x40, 0xad, 0xe6, 0x06, 0x8d, 0xfe, 0x2f, 0x44, 0xe8, 0xfa, 0x73, 0x3c},
			[32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xd3, 0xc2, 0x1b, 0xce, 0xcc, 0xed, 0xca, 0x35, 0xb7, 0xd9, 0x15, 0x45, 0x8e, 0xf5, 0x40, 0xad, 0xe6, 0x06, 0x8d, 0xfe, 0x2f, 0x44, 0xe8, 0xfa, 0x73, 0x3c},
			[32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x1e, 0x19, 0xe0, 0xc9, 0xba, 0xca, 0x35, 0xb7, 0xd9, 0x15, 0x45, 0x8e, 0xf5, 0x40, 0xad, 0xe6, 0x06, 0x8d, 0xfe, 0x2f, 0x44, 0xe8, 0xfa, 0x73, 0x3c},
			[32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x8c, 0xf2, 0x3f, 0x90, 0x9c, 0xca, 0x35, 0xb7, 0xd9, 0x15, 0x45, 0x8e, 0xf5, 0x40, 0xad, 0xe6, 0x06, 0x8d, 0xfe, 0x2f, 0x44, 0xe8, 0xfa, 0x73, 0x3c},
			[32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x92, 0x5e, 0x06, 0xee, 0xc9, 0xca, 0x35, 0xb7, 0xd9, 0x15, 0x45, 0x8e, 0xf5, 0x40, 0xad, 0xe6, 0x06, 0x8d, 0xfe, 0x2f, 0x44, 0xe8, 0xfa, 0x73, 0x3c},
			[32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x97, 0xc9, 0xce, 0x4c, 0xf6, 0xca, 0x35, 0xb7, 0xd9, 0x15, 0x45, 0x8e, 0xf5, 0x40, 0xad, 0xe6, 0x06, 0x8d, 0xfe, 0x2f, 0x44, 0xe8, 0xfa, 0x73, 0x3c},
			[32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x9d, 0x43, 0x76, 0x61, 0xd7, 0xca, 0x35, 0xb7, 0xd9, 0x15, 0x45, 0x8e, 0xf5, 0x40, 0xad, 0xe6, 0x06, 0x8d, 0xfe, 0x2f, 0x44, 0xe8, 0xfa, 0x73, 0x3c},
		},
		pubkeys: [][32]byte{
			[32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01},
			[32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02},
			[32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03},
			[32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04},
			[32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05},
			[32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x06},
			[32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x07},
			[32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08},
			[32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x09},
		},
	},
	{
		compvals: [][32]byte{
			[32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xa2, 0xa1, 0x5d, 0x09, 0x51, 0xca, 0x35, 0xb7, 0xd9, 0x15, 0x45, 0x8e, 0xf5, 0x40, 0xad, 0xe6, 0x06, 0x8d, 0xfe, 0x2f, 0x44, 0xe8, 0xfa, 0x73, 0x3c},
			[32]byte{0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x87, 0x86, 0x78, 0x32, 0x6e, 0xca, 0x35, 0xb7, 0xd9, 0x15, 0x45, 0x8e, 0xf5, 0x40, 0xad, 0xe6, 0x06, 0x8d, 0xfe, 0x2f, 0x44, 0xe8, 0xfa, 0x73, 0x3c},
			[32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x45, 0x42, 0xba, 0x12, 0xa3, 0xca, 0x35, 0xb7, 0xd9, 0x15, 0x45, 0x8e, 0xf5, 0x40, 0xad, 0xe6, 0x06, 0x8d, 0xfe, 0x2f, 0x44, 0xe8, 0xfa, 0x73, 0x3c},
			[32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xd3, 0xc2, 0x1b, 0xce, 0xcc, 0xed, 0xca, 0x35, 0xb7, 0xd9, 0x15, 0x45, 0x8e, 0xf5, 0x40, 0xad, 0xe6, 0x06, 0x8d, 0xfe, 0x2f, 0x44, 0xe8, 0xfa, 0x73, 0x3c},
			[32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x1e, 0x19, 0xe0, 0xc9, 0xba, 0xca, 0x35, 0xb7, 0xd9, 0x15, 0x45, 0x8e, 0xf5, 0x40, 0xad, 0xe6, 0x06, 0x8d, 0xfe, 0x2f, 0x44, 0xe8, 0xfa, 0x73, 0x3c},
			[32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x8c, 0xf2, 0x3f, 0x90, 0x9c, 0xca, 0x35, 0xb7, 0xd9, 0x15, 0x45, 0x8e, 0xf5, 0x40, 0xad, 0xe6, 0x06, 0x8d, 0xfe, 0x2f, 0x44, 0xe8, 0xfa, 0x73, 0x3c},
			[32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x95, 0x5e, 0x06, 0xee, 0xc9, 0xca, 0x35, 0xb7, 0xd9, 0x15, 0x45, 0x8e, 0xf5, 0x40, 0xad, 0xe6, 0x06, 0x8d, 0xfe, 0x2f, 0x44, 0xe8, 0xfa, 0x73, 0x3c},
			[32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x97, 0xc9, 0xce, 0x4c, 0xf6, 0xca, 0x35, 0xb7, 0xd9, 0x15, 0x45, 0x8e, 0xf5, 0x40, 0xad, 0xe6, 0x06, 0x8d, 0xfe, 0x2f, 0x44, 0xe8, 0xfa, 0x73, 0x3c},
			[32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x9d, 0x43, 0x76, 0x61, 0xd7, 0xca, 0x35, 0xb7, 0xd9, 0x15, 0x45, 0x8e, 0xf5, 0x40, 0xad, 0xe6, 0x06, 0x8d, 0xfe, 0x2f, 0x44, 0xe8, 0xfa, 0x73, 0x3c},
		},
		pubkeys: [][32]byte{
			[32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01},
			[32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02},
			[32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03},
			[32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a},
			[32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05},
			[32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x06},
			[32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x07},
			[32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0b},
			[32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x09},
		},
	},
}

func Test_getNewValidators(t *testing.T) {
	type args struct {
		compvals [][32]byte
		pubkeys  [][32]byte
	}
	tests := []struct {
		name  string
		args  args
		want  []abciTypes.Validator
		want1 []byte
	}{
		{
			name: "Test1",
			args: args{
				compvals: testData[0].compvals,
				pubkeys:  testData[0].pubkeys,
			},
			want1: []byte{49, 127, 204, 41, 64, 181, 120, 110, 28, 12, 152, 151, 13, 9, 15, 87, 209, 246, 237, 154, 101, 210, 164, 153, 117, 241, 235, 78, 214, 255, 72, 36},
		},
		{
			name: "Test2",
			args: args{
				compvals: testData[1].compvals,
				pubkeys:  testData[1].pubkeys,
			},
			want1: []byte{179, 235, 155, 247, 195, 155, 123, 3, 84, 116, 142, 170, 61, 10, 222, 112, 56, 5, 33, 76, 123, 76, 119, 123, 87, 38, 79, 97, 216, 245, 224, 137},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := getNewValidators(tt.args.compvals, tt.args.pubkeys)
			//We'll compare hashes only
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("getNewValidators() got = %v, want %v", got, tt.want)
			//}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("getNewValidators() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_getUpdatedValidators(t *testing.T) {
	type args struct {
		newValidators []abciTypes.Validator
		oldValidators []abciTypes.Validator
	}

	oldValidators, _ := getNewValidators(testData[1].compvals, testData[1].pubkeys)
	newValidators, _ := getNewValidators(testData[2].compvals, testData[2].pubkeys)

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test1",
			args: args{
				newValidators: newValidators,
				oldValidators: oldValidators,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getUpdatedValidators(tt.args.newValidators, tt.args.oldValidators); len(got) != tt.want {
				t.Errorf("getUpdatedValidators() = %v, want %v", got, tt.want)
			}
		})
	}
}