package numbers

import "testing"

func Test_isStraight(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "有大小王",
			args: args{
				[]int{0, 0, 1, 2, 5},
			},
			want: true,
		},
		{
			name: "没有大小王",
			args: args{
				[]int{1, 2, 4, 5, 6},
			},
			want: false,
		},
		{
			name: "有小王或大王",
			args: args{
				[]int{0, 2, 4, 5, 6},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStraight(tt.args.nums); got != tt.want {
				t.Errorf("isStraight() = %v, want %v", got, tt.want)
			}
		})
	}
}
