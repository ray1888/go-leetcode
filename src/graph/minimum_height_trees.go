package graph

import (
	"fmt"
	"math"
)

func dfs(graph [][]int, node int, visited []bool, dist int, max *int) {
	visited[node] = true
	if dist > *max {
		*max = dist
	}
	for _, item := range graph[node] {
		if !visited[item] {
			dfs(graph, item, visited, dist+1, max)
		}
	}
	visited[node] = false
}

func findMinHeightTreesBrutualForce(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, 0)
	}

	for _, item := range edges {
		graph[item[0]] = append(graph[item[0]], item[1])
		graph[item[1]] = append(graph[item[1]], item[0])
	}
	fmt.Println(graph)

	visited := make([]bool, n)
	height := make([]int, n)
	max := 0
	minHeight := math.MaxInt32
	for i := 0; i < n; i++ {
		dfs(graph, i, visited, 0, &max)
		height[i] = max
		if minHeight > height[i] {
			minHeight = height[i]
		}
		max = 0
	}
	var result []int
	for index, value := range height {
		if value == minHeight {
			result = append(result, index)
		}
	}
	return result
}

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, 0)
	}

	for _, item := range edges {
		graph[item[0]] = append(graph[item[0]], item[1])
		graph[item[1]] = append(graph[item[1]], item[0])
	}

	degree := make([]int, n)
	leaves := make([]int, 0)
	for node, items := range graph {
		degree[node] = len(items)
		if degree[node] == 1 {
			leaves = append(leaves, node)
		}
	}

	for len(leaves) < n {
		size := len(leaves)
		n -= size
		for i := 0; i < size; i++ {
			leaf := leaves[0]
			leaves = leaves[1:]
			for _, item := range graph[leaf] {
				degree[item]--
				if degree[item] == 1 {
					leaves = append(leaves, item)
				}
			}
		}
	}
	return leaves
}
