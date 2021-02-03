package backtracing

import "testing"

func Test_getPermutation2(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				n: 3,
				k: 3,
			},
			want: "213",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPermutation2(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("getPermutation2() = %v, want %v", got, tt.want)
			}
		})
	}
}
