package stock

import "testing"

func Test_maxProfitCoolDown(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				prices: []int{1, 2, 3, 0, 2},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitCoolDown(tt.args.prices); got != tt.want {
				t.Errorf("maxProfitCoolDown() = %v, want %v", got, tt.want)
			}
		})
	}
}
