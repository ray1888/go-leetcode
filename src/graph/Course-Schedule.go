package graph

func hasCycle(graph [][]int, checked []bool, visited []bool, index int) bool {
	if visited[index] {
		return true
	}
	visited[index] = true
	for _, item := range graph[index] {
		if !checked[item] && hasCycle(graph, checked, visited, item) {
			return true
		}
	}
	visited[index] = false
	checked[index] = true
	return false
}

func canFinishDFS(numCourses int, prerequisites [][]int) bool {
	if numCourses == 0 {
		return true
	}
	if len(prerequisites) == 0 {
		return true
	}

	graph := make([][]int, numCourses*numCourses)
	for i := 0; i <= numCourses; i++ {
		graph[i] = make([]int, 0)
	}
	visited := make([]bool, numCourses)
	checked := make([]bool, numCourses)

	for _, item := range prerequisites {
		graph[item[1]] = append(graph[item[1]], item[0])
	}

	for i := 0; i < numCourses; i++ {
		if !checked[i] && hasCycle(graph, checked, visited, i) {
			return false
		}
	}
	return true
}

func canFinishTOPOSort(numCourses int, prerequisites [][]int) bool {
	if numCourses == 0 {
		return true
	}
	if len(prerequisites) == 0 {
		return true
	}

	graphTable := make([][]int, numCourses*numCourses)
	inDegree := make([]int, numCourses)

	for i := 0; i < numCourses; i++ {
		graphTable[i] = make([]int, 0, numCourses)
	}

	for _, i := range prerequisites {
		after, before := i[0], i[1]
		graphTable[before] = append(graphTable[before], after)
		inDegree[after] += 1
	}

	Queue := make([]int, 0, numCourses)

	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			Queue = append(Queue, i)
		}
	}

	count := 0
	for len(Queue) > 0 {
		v := Queue[0]
		Queue = Queue[1:]
		count += 1
		nextNodeList := graphTable[v]
		for _, i := range nextNodeList {
			inDegree[i] -= 1
			if inDegree[i] == 0 {
				Queue = append(Queue, i)
			}
		}

	}
	return count == numCourses
}

// dfs Map 版本
// 如果使用Map版本，进阶版需要返回顺序的情况是不能处理的
/*
	题解：
	1. dfs回溯算法
	2. 有一个十分重要的条件，当某个点不成环的情况下，如果这个点为下一个节点，则上一个节点如果有且只有这么一个点的情况下，也是不成环的
*/
func hasCycleMap(graph map[int][]int, visitMap map[int]bool, checkMap map[int]bool, index int) bool {
	if visitMap[index] {
		return true
	}
	visitMap[index] = true
	for _, item := range graph[index] {
		if !checkMap[item] && hasCycleMap(graph, visitMap, checkMap, item) {
			return true
		}
	}
	visitMap[index] = false
	checkMap[index] = true
	return false
}

func canFinishDFSMap(numCourses int, prerequisites [][]int) bool {
	if numCourses == 0 || len(prerequisites) == 0 {
		return true
	}
	visitMap := make(map[int]bool)
	checkMap := make(map[int]bool)
	graph := make(map[int][]int)

	for _, item := range prerequisites {
		if _, ok := graph[item[1]]; !ok {
			graph[item[1]] = make([]int, 0)
		}
		if _, ok := visitMap[item[1]]; !ok {
			visitMap[item[1]] = false
		}
		if _, ok := checkMap[item[1]]; !ok {
			checkMap[item[1]] = false
		}
		graph[item[1]] = append(graph[item[1]], item[0])
	}

	for i := 0; i < numCourses; i++ {
		if !checkMap[i] && hasCycleMap(graph, visitMap, checkMap, i) {
			return false
		}
	}
	return true

}
