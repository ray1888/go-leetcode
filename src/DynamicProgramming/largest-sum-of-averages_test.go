package DynamicProgramming

import "testing"

func Test_largestSumOfAveragesSpaceOn(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			args: args{
				A: []int{9, 1, 2, 3, 9},
				K: 3,
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestSumOfAveragesspaceOn(tt.args.A, tt.args.K); got != tt.want {
				t.Errorf("largestSumOfAveragesSpaceOnSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_largestSumOfAveragesSpaceOnSquare(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			args: args{
				A: []int{9, 1, 2, 3, 9},
				K: 3,
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestSumOfAveragesSpaceOnSquare(tt.args.A, tt.args.K); got != tt.want {
				t.Errorf("largestSumOfAveragesSpaceOnSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
