package DynamicProgramming

import (
	"testing"
)

func Test_backPackIII(t *testing.T) {
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
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backPackIII(tt.args.m, tt.args.A, tt.args.V); got != tt.want {
				t.Errorf("backPackIII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backPackIIIOnSpace(t *testing.T) {
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
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backPackIIIOnSpace(tt.args.m, tt.args.A, tt.args.V); got != tt.want {
				t.Errorf("backPackIIIOnSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
