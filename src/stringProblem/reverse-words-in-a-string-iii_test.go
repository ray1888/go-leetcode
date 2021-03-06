package stringProblem

import "testing"

func Test_reverseWordsiii(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s: "Let's take LeetCode contest",
			},
			want: "s'teL ekat edoCteeL tsetnoc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWordsiii(tt.args.s); got != tt.want {
				t.Errorf("reverseWordsiii() = %v, want %v", got, tt.want)
			}
		})
	}
}
