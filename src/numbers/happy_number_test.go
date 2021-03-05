package numbers

import (
	"testing"
)

func TestHappyNumberTwoPointer(t *testing.T) {
	tests := []BoolTest{
		{Name: "is happy number", Input: 28, Output: true},
		{Name: "is not happy number", Input: 2, Output: false},
	}
	for _, test := range tests {
		result := HappyNubmerTwoPointer(test.Input)
		if result != test.Output {
			t.Fatalf("Test %s failed", test.Name)
		}

	}
}

func Test_isHappy(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				n: 19,
			},
			want: true,
		},
		{
			args: args{
				n: 2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy(tt.args.n); got != tt.want {
				t.Errorf("isHappy() = %v, want %v", got, tt.want)
			}
		})
	}
}
