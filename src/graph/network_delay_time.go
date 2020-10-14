package graph

import (
	"math"
	"sort"
)

func minDistanceIndex(st []bool, dist []int) int {
	var minIndex int
	var minVal = math.MaxInt32
	for i := 1; i < len(dist); i++ {
		if !st[i] && dist[i] <= minVal {
			minIndex = i
			minVal = dist[i]
		}
	}
	return minIndex
}

func networkDelayTime(times [][]int, N int, K int) int {
	graph := make([][]int, N+1)
	for index := range graph {
		graph[index] = make([]int, N+1)
		for k := 0; k < N+1; k++ {
			graph[index][k] = -1
		}
	}
	dist := make([]int, N+1)
	for index := range dist {
		dist[index] = math.MaxInt32
	}

	dist[K] = 0
	for _, item := range times {
		graph[item[0]][item[1]] = item[2]
	}
	st := make([]bool, N+1)
	for i := 1; i <= N; i++ {
		index := minDistanceIndex(st, dist)
		st[index] = true
		for item := range graph[index] {
			if !st[item] && graph[index][item] != -1 && dist[index] != math.MaxInt32 && dist[item] > dist[index]+graph[index][item] {
				dist[item] = dist[index] + graph[index][item]
			}
		}
	}
	dist = dist[1:]
	sort.Ints(dist)
	if dist[N-1] == math.MaxInt32 {
		return -1
	}
	return dist[N-1]
}
