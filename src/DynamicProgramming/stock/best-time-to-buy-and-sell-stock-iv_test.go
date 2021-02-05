package stock

import "testing"

func Test_maxProfitIV(t *testing.T) {
	type args struct {
		K      int
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				K:      2,
				prices: []int{2, 4, 1},
			},
			want: 2,
		},
		{
			args: args{
				K:      2,
				prices: []int{3, 2, 6, 5, 0, 3},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitIV(tt.args.K, tt.args.prices); got != tt.want {
				t.Errorf("maxProfitIV() = %v, want %v", got, tt.want)
			}
		})
	}
}
