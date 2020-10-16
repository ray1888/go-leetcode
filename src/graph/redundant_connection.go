package graph

import (
	"github.com/golang-collections/collections/set"
)

func hasCycleRC(cur int, target int, graph map[int]*set.Set, pre int) bool {
	if graph[cur].Has(target) {
		return true
	}
	graph[cur].Do()
	for _, num := range  {
		if num == pre {
			continue
		}
	}
	return false
}

func findRedundantConnection(edges [][]int) []int {
	graph := make(map[int]*set.Set)
	for _, edge := range edges {
		if _, ok := graph[edge[0]]; !ok {
			graph[edge[0]] = set.New()
		}
		if _, ok := graph[edge[1]]; !ok {
			graph[edge[1]] = set.New()
		}
		graph[edge[0]].Insert(edge[1])
		graph[edge[1]].Insert(edge[0])
	}

	for _, edge := range edges {
		graph[edge[0]].Remove(edge[1])
		graph[edge[1]].Remove(edge[0])
		if hasCycleRC(edge[0], edge[1], graph, -1) {
			return edge
		} else {
			graph[edge[0]].Insert(edge[1])
			graph[edge[1]].Insert(edge[0])
		}
	}
	return []int{}
}
