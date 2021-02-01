package DynamicProgramming

import "testing"

func Test_maxSubArrayLen(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{1, -1, 5, -2, 3},
				k:    3,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArrayLen(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maxSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
