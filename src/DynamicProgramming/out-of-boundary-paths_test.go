package DynamicProgramming

import "testing"

func Test_findPaths(t *testing.T) {
	type args struct {
		m int
		n int
		N int
		i int
		j int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{2, 2, 2, 0, 0},
			want: 6,
		},

		{
			args: args{1, 3, 3, 0, 1},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPaths(tt.args.m, tt.args.n, tt.args.N, tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("findPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
