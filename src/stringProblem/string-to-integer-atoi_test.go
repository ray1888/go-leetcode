package stringProblem

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				str: "42",
			},
			want: 42,
		},
		{
			args: args{
				str: "-92",
			},
			want: -92,
		},
		{
			args: args{
				str: "+-12",
			},
			want: 0,
		},
		{
			args: args{
				str: "-+12",
			},
			want: 0,
		},
		{
			args: args{
				str: "-91283472332",
			},
			want: -2147483648,
		},
		{
			args: args{
				str: "words and 987",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
