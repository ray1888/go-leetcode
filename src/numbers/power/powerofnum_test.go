package power

import "testing"

func Test_isPowerOfK(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "power of two",
			args: args{
				n: 4,
				k: 2,
			},
			want: true,
		},
		{
			name: "power of two-2",
			args: args{
				n: 5,
				k: 2,
			},
			want: false,
		},
		{
			name: "power of three",
			args: args{
				n: 9,
				k: 3,
			},
			want: true,
		},
		{
			name: "power of three-2",
			args: args{
				n: 8,
				k: 3,
			},
			want: false,
		},
		{
			name: "power of four",
			args: args{
				n: 16,
				k: 4,
			},
			want: true,
		},
		{
			name: "power of four-2",
			args: args{
				n: 5,
				k: 4,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfK(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("isPowerOfK() = %v, want %v", got, tt.want)
			}
		})
	}
}
