package numbers

import (
	"reflect"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				matrix: [][]int{
					[]int{1, 2, 3},
					[]int{4, 5, 6},
					[]int{7, 8, 9},
				},
			},
			want: []int{
				1, 2, 3, 6, 9, 8, 7, 4, 5,
			},
		},
		{
			args: args{
				matrix: [][]int{
					[]int{1, 2, 3, 4},
					[]int{5, 6, 7, 8},
					[]int{9, 10, 11, 12},
				},
			},
			want: []int{
				1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralOrder(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
