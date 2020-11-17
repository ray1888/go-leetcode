package DynamicProgramming

import "testing"

func TestMinCostWithKColors(t *testing.T) {
	type args struct {
		nums [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums: [][]int{
					[]int{8, 2, 4},
					[]int{5, 7, 3},
					[]int{9, 1, 6},
					[]int{4, 1, 9},
				},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinCostWithKColors(tt.args.nums); got != tt.want {
				t.Errorf("MinCostWithKColors() = %v, want %v", got, tt.want)
			}
		})
	}
}
