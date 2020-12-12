package DynamicProgramming

import (
	"testing"
)

func Test_backPackII(t *testing.T) {
	type args struct {
		m int
		A []int
		V []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				m: 10,
				A: []int{2, 3, 5, 7},
				V: []int{1, 5, 2, 4},
			},
			want: 9,
		},
		{
			args: args{
				m: 10,
				A: []int{2, 3, 8},
				V: []int{2, 5, 8},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backPackII(tt.args.m, tt.args.A, tt.args.V); got != tt.want {
				t.Errorf("backPackII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backPackIIOnSpace(t *testing.T) {
	type args struct {
		m int
		A []int
		V []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				m: 10,
				A: []int{2, 3, 5, 7},
				V: []int{1, 5, 2, 4},
			},
			want: 9,
		},
		{
			args: args{
				m: 10,
				A: []int{2, 3, 8},
				V: []int{2, 5, 8},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backPackIIOnSpace(tt.args.m, tt.args.A, tt.args.V); got != tt.want {
				t.Errorf("backPackIIOnSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
