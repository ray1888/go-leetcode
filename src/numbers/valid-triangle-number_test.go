package numbers

import "testing"

func Test_triangleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{2, 2, 3, 4},
			},
			want: 3,
		},
		{
			args: args{
				nums: []int{0, 0, 0},
			},
			want: 0,
		},
		{
			args: args{
				nums: []int{0, 1, 0, 1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := triangleNumber(tt.args.nums); got != tt.want {
				t.Errorf("triangleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
