package numbers

import "testing"

func Test_nextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args:{
				nums:[]int{1,2,3},
			},
			want: []int{1,3,2},
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.args.nums)
			t.assertEqual(tt.args.nums, tt.want)
		})
	}
}
