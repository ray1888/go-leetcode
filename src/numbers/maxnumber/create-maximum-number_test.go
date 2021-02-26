package maxnumber

import (
	"reflect"
	"testing"
)

func Test_maxNumber(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		k     int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{
			args: args{
				nums1: []int{3, 4, 6, 5},
				nums2: []int{9, 1, 2, 5, 8, 3},
				k:     5,
			},
			wantRes: []int{9, 8, 6, 5, 3},
		},
		{
			args: args{
				nums1: []int{6, 7},
				nums2: []int{6, 0, 4},
				k:     5,
			},
			wantRes: []int{6, 7, 6, 0, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := maxNumber(tt.args.nums1, tt.args.nums2, tt.args.k); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("maxNumber() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
