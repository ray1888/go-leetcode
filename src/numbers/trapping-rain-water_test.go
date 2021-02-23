package numbers

import (
	"testing"
)

func Test_trapSpaceOn(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trapSpaceOn(tt.args.height); got != tt.want {
				t.Errorf("trapSpaceOn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trapSpaceO1(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trapSpaceO1(tt.args.height); got != tt.want {
				t.Errorf("trapSpaceO1() = %v, want %v", got, tt.want)
			}
		})
	}
}
