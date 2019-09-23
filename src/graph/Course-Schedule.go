package graph

func hasCycle(graph [][]int, visited, checked []bool, v int) bool {
	if visited[v] {
		return true
	}
	visited[v] = true
	for _, index := range graph[v] {
		if !checked[index] && hasCycle(graph, visited, checked, index) {
			return true
		}
	}
	checked[v] = true
	visited[v] = false
	return false
}

func canFinishDFS(numCourses int, prerequisites [][]int) bool {
	if numCourses == 0 {
		return true
	}
	if len(prerequisites) == 0 {
		return true
	}

	graphTable := make([][]int, numCourses*numCourses)

	for i := 0; i < numCourses; i++ {
		graphTable[i] = make([]int, 0, numCourses)
	}

	for _, i := range prerequisites {
		after, before := i[0], i[1]
		graphTable[before] = append(graphTable[before], after)
	}

	checked := make([]bool, numCourses)
	visited := make([]bool, numCourses)

	for i := 0; i < numCourses; i++ {
		if !checked[i] && hasCycle(graphTable, visited, checked, i) {
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
