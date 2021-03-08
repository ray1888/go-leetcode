package montonicstack

import (
	"reflect"
	"testing"
)

func Test_dailyTemperatures(t *testing.T) {
	type args struct {
		T []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				T: []int{73, 74, 75, 71, 69, 72, 76, 73},
			},
			want: []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dailyTemperatures(tt.args.T); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dailyTemperatures() = %v, want %v", got, tt.want)
			}
		})
	}
}
