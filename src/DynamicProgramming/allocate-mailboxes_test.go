package DynamicProgramming

import "testing"

func Test_minDistanceMailBox(t *testing.T) {
	type args struct {
		houses []int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				houses: []int{1, 4, 8, 10, 20},
				k:      3,
			},
			want: 5,
		},
		{
			args: args{
				houses: []int{7, 4, 6, 1},
				k:      1,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistanceMailBox(tt.args.houses, tt.args.k); got != tt.want {
				t.Errorf("minDistanceMailBox() = %v, want %v", got, tt.want)
			}
		})
	}
}
