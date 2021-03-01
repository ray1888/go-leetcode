package numbers

import (
	"testing"
)

func TestSplitBrutalForce(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				arr: []int{4, 7, 2, 10, 5},
				k:   5,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitBrutalForce(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("SplitBrutalForce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitHalfCut(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				arr: []int{4, 7, 2, 10, 5},
				k:   5,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitHalfCut(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("splitHalfCut() = %v, want %v", got, tt.want)
			}
		})
	}
}
