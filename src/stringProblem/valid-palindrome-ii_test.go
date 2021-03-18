package stringProblem

import "testing"

func Test_validPalindrome2(t *testing.T) {
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
				s: "aba",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPalindrome2(tt.args.s); got != tt.want {
				t.Errorf("validPalindrome2() = %v, want %v", got, tt.want)
			}
		})
	}
}
