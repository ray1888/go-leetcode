package binary_search

import "testing"

func TestFindRepeatCountOLogN(t *testing.T) {
	type args struct {
		arr       []int
		repeatNum int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{

				arr:       []int{1, 2, 2, 2, 5, 6},
				repeatNum: 2,
			},
			want: 3,
		},
		{
			args: args{

				arr:       []int{1, 2, 3, 5, 5},
				repeatNum: 5,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindRepeatCountOLogN(tt.args.arr, tt.args.repeatNum); got != tt.want {
				t.Errorf("FindRepeatCountOLogN() = %v, want %v", got, tt.want)
			}
		})
	}
}
