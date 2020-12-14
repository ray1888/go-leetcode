package DynamicProgramming

import "testing"

func Test_integerBreak(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			// 10 = 3+ 3+ 4 , 3 * 3 * 4 =36
			args: args{
				n: 10,
			},
			want: 36,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := integerBreak(tt.args.n); got != tt.want {
				t.Errorf("integerBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
