package backtracing

import (
	"reflect"
	"testing"
)

func Test_restoreIpAddresses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// {
		// 	name: "normal",
		// 	args: args{
		// 		s: "25525511135",
		// 	},
		// 	want: []string{"255.255.11.135", "255.255.111.35"},
		// },
		{
			name: "zero expect",
			args: args{
				s: "010010",
			},
			want: []string{"0.10.0.10", "0.100.1.0"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreIpAddresses(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreIpAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}
