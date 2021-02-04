package numbers

import "testing"

func Test_validMountainArray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				arr: []int{0, 3, 2, 1},
			},
			want: true,
		},
		{
			args: args{
				arr: []int{3, 5, 5},
			},
			want: false,
		},
		{
			args: args{
				arr: []int{2, 1},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validMountainArray(tt.args.arr); got != tt.want {
				t.Errorf("validMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
