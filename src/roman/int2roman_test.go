package roman

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				num: 4,
			},
			want: "IV",
		},
		{
			args: args{
				num: 3,
			},
			want: "III",
		},
		{
			args: args{
				num: 58,
			},
			want: "LVIII",
		},
		{
			args: args{
				num: 1994,
			},
			want: "MCMXCIV",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
