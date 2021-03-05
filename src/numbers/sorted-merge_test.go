package numbers

import "testing"

func Test_merge(t *testing.T) {
	type args struct {
		A []int
		m int
		B []int
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				A: []int{1, 2, 3, 0, 0, 0},
				B: []int{2, 5, 6},
				m: 3,
				n: 3,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.A, tt.args.m, tt.args.B, tt.args.n)
		})
	}
}
