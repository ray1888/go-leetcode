package graph

import (
	"fmt"
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

func BellmanFord(times [][]int, N int, K int) int {
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
	// 松弛算法，对每个可达得边全部计算一次
	for i := 0; i < N-1; i++ {
		for j := 0; j < len(times); j++ {
			u := times[j][0]
			v := times[j][1]
			weight := times[j][2]
			if dist[u] != math.MaxInt32 && dist[v] > dist[u]+weight {
				dist[v] = dist[u] + weight
			}
		}
	}
	// 计算是否含有负环
	for i := 0; i < len(times); i++ {
		u := times[i][0]
		v := times[i][1]
		weight := times[i][2]
		if dist[u] != math.MaxInt32 && dist[v] > dist[u]+weight {
			fmt.Println("has negative cycle")
			return -1
		}
	}
	dist = dist[1:]
	sort.Ints(dist)
	if dist[N-1] == math.MaxInt32 {
		return -1
	}
	return dist[N-1]
}

func SPFAAlogorithm(times [][]int, N int, K int) int {
	graph := make([][]int, N+1)
	for index := range graph {
		graph[index] = make([]int, N+1)
		for k := 0; k < N+1; k++ {
			graph[index][k] = math.MaxInt32
		}
	}
	dist := make([]int, N+1)
	WorkingQueue := make([]int, 0)
	InQueue := make([]bool, N+1)
	for index := range dist {
		dist[index] = math.MaxInt32
	}
	for _, item := range times {
		graph[item[0]][item[1]] = item[2]
	}
	dist[K] = 0
	InQueue[K] = true
	WorkingQueue = append(WorkingQueue, K)
	for len(WorkingQueue) > 0 {
		u := WorkingQueue[0]
		WorkingQueue = WorkingQueue[1:]
		InQueue[u] = false
		for v, val := range graph[u] {
			// if val not in queue, append
			if v == 0 || v == u {
				continue
			}
			if dist[v] > dist[u]+val {
				dist[v] = dist[u] + val
				if !InQueue[v] {
					WorkingQueue = append(WorkingQueue, v)
					InQueue[v] = true
				}
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

func FloydWarshallAlgorithm(times [][]int, N int, K int) int {
	graph := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		emptyList := make([]int, N+1)
		for j := 0; j < N+1; j++ {
			emptyList[j] = math.MaxInt32
		}
		graph[i] = emptyList
	}

	dist := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		emptyList := make([]int, N+1)
		for j := 0; j < N+1; j++ {
			emptyList[j] = math.MaxInt32
		}
		dist[i] = emptyList
	}

	for _, item := range times {
		graph[item[0]][item[1]] = item[2]
		dist[item[0]][item[1]] = item[2]
	}

	for k := 1; k < N+1; k++ {
		for i := 1; i < N+1; i++ {
			for j := 1; j < N+1; j++ {
				if dist[i][j] > dist[i][k]+dist[k][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}
	sort.Ints(dist[K])
	maxValue := 0
	for _, item := range dist[K] {
		if item != math.MaxInt32 && item > maxValue {
			maxValue = item
		}
	}
	return maxValue
}
