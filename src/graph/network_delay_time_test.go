package graph

import "testing"

func Test_networkDelayTime(t *testing.T) {
	type args struct {
		times [][]int
		N     int
		K     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal",
			args: args{
				times: [][]int{
					{2, 1, 1},
					{2, 3, 1},
					{3, 4, 1},
				},
				N: 4,
				K: 2,
			},
			want: 2,
		},
		{
			name: "extra long",
			args: args{
				times: [][]int{
					{3, 5, 78},
					{2, 1, 1},
					{1, 3, 0},
					{4, 3, 59},
					{5, 3, 85},
					{5, 2, 22},
					{2, 4, 23},
					{1, 4, 43},
					{4, 5, 75},
					{5, 1, 15},
					{1, 5, 91},
					{4, 1, 16},
					{3, 2, 98},
					{3, 4, 22},
					{5, 4, 31},
					{1, 2, 0},
					{2, 5, 4},
					{4, 2, 51},
					{3, 1, 36},
					{2, 3, 59},
				},
				N: 5,
				K: 5,
			},
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := networkDelayTime(tt.args.times, tt.args.N, tt.args.K); got != tt.want {
				t.Errorf("networkDelayTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_BellmendFord(t *testing.T) {
	type args struct {
		times [][]int
		N     int
		K     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal",
			args: args{
				times: [][]int{
					{2, 1, 1},
					{2, 3, 1},
					{3, 4, 1},
				},
				N: 4,
				K: 2,
			},
			want: 2,
		},
		{
			name: "extra long",
			args: args{
				times: [][]int{
					{3, 5, 78},
					{2, 1, 1},
					{1, 3, 0},
					{4, 3, 59},
					{5, 3, 85},
					{5, 2, 22},
					{2, 4, 23},
					{1, 4, 43},
					{4, 5, 75},
					{5, 1, 15},
					{1, 5, 91},
					{4, 1, 16},
					{3, 2, 98},
					{3, 4, 22},
					{5, 4, 31},
					{1, 2, 0},
					{2, 5, 4},
					{4, 2, 51},
					{3, 1, 36},
					{2, 3, 59},
				},
				N: 5,
				K: 5,
			},
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BellmanFord(tt.args.times, tt.args.N, tt.args.K); got != tt.want {
				t.Errorf("networkDelayTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_SPFAAlgorithm(t *testing.T) {
	type args struct {
		times [][]int
		N     int
		K     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal",
			args: args{
				times: [][]int{
					{2, 1, 1},
					{2, 3, 1},
					{3, 4, 1},
				},
				N: 4,
				K: 2,
			},
			want: 2,
		},
		{
			name: "extra long",
			args: args{
				times: [][]int{
					{3, 5, 78},
					{2, 1, 1},
					{1, 3, 0},
					{4, 3, 59},
					{5, 3, 85},
					{5, 2, 22},
					{2, 4, 23},
					{1, 4, 43},
					{4, 5, 75},
					{5, 1, 15},
					{1, 5, 91},
					{4, 1, 16},
					{3, 2, 98},
					{3, 4, 22},
					{5, 4, 31},
					{1, 2, 0},
					{2, 5, 4},
					{4, 2, 51},
					{3, 1, 36},
					{2, 3, 59},
				},
				N: 5,
				K: 5,
			},
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SPFAAlogorithm(tt.args.times, tt.args.N, tt.args.K); got != tt.want {
				t.Errorf("networkDelayTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_FloydWarshallAlgorithm(t *testing.T) {
	type args struct {
		times [][]int
		N     int
		K     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal",
			args: args{
				times: [][]int{
					{2, 1, 1},
					{2, 3, 1},
					{3, 4, 1},
				},
				N: 4,
				K: 2,
			},
			want: 2,
		},
		{
			name: "extra long",
			args: args{
				times: [][]int{
					{3, 5, 78},
					{2, 1, 1},
					{1, 3, 0},
					{4, 3, 59},
					{5, 3, 85},
					{5, 2, 22},
					{2, 4, 23},
					{1, 4, 43},
					{4, 5, 75},
					{5, 1, 15},
					{1, 5, 91},
					{4, 1, 16},
					{3, 2, 98},
					{3, 4, 22},
					{5, 4, 31},
					{1, 2, 0},
					{2, 5, 4},
					{4, 2, 51},
					{3, 1, 36},
					{2, 3, 59},
				},
				N: 5,
				K: 5,
			},
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloydWarshallAlgorithm(tt.args.times, tt.args.N, tt.args.K); got != tt.want {
				t.Errorf("networkDelayTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
