package statemachine

import "testing"

func Test_isNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				s: "+100",
			},
			want: true,
		},
		{
			args: args{
				s: "5e2",
			},
			want: true,
		},
		{
			args: args{
				s: "3.14",
			},
			want: true,
		},
		{
			args: args{
				s: "1a3.14",
			},
			want: false,
		},
		{
			args: args{
				s: "+-5",
			},
			want: false,
		},
		{
			args: args{
				s: "12e+5.4",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumber(tt.args.s); got != tt.want {
				t.Errorf("isNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
