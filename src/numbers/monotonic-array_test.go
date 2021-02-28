package numbers

import (
	"testing"
)

func Test_isMonotonic(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				A: []int{1, 2, 3, 3},
			},
			want: true,
		},
		{
			args: args{
				A: []int{6, 5, 4, 4},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMonotonic(tt.args.A); got != tt.want {
				t.Errorf("isMonotonic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isMonotonicOn(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				A: []int{1, 2, 3, 3},
			},
			want: true,
		},
		{
			args: args{
				A: []int{6, 5, 4, 4},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMonotonicOn(tt.args.A); got != tt.want {
				t.Errorf("isMonotonicOn() = %v, want %v", got, tt.want)
			}
		})
	}
}
