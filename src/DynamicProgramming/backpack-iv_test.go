package DynamicProgramming

import "testing"

func Test_backPackIV(t *testing.T) {
	type args struct {
		m     int
		A     []int
		V     []int
		Times []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				m:     10,
				A:     []int{2, 3, 5, 7},
				V:     []int{1, 5, 2, 4},
				Times: []int{0, 0, 0, 0},
			},
			want: 15,
		},
		{
			args: args{
				m:     10,
				A:     []int{2, 3, 5, 7},
				V:     []int{1, 5, 2, 4},
				Times: []int{0, 2, 0, 0},
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backPackIV(tt.args.m, tt.args.A, tt.args.V, tt.args.Times); got != tt.want {
				t.Errorf("backPackIV() = %v, want %v", got, tt.want)
			}
		})
	}
}
