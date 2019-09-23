package graph

func hasCycleFindOrder(graph [][]int, visited, checked []bool, result *[]int, v int) bool {
	if visited[v] {
		return true
	}
	visited[v] = true
	for _, index := range graph[v] {
		if !checked[index] && hasCycleFindOrder(graph, visited, checked, result, index) {
			return true
		}
	}
	*result = append(*result, v)
	checked[v] = true
	visited[v] = false
	return false
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	if numCourses == 0 {
		return []int{}
	}

	if len(prerequisites) == 0 {
		result := make([]int, numCourses)
		for i := 0; i < numCourses; i++ {
			result[numCourses-i-1] = i
		}
		return result
	}

	graphTable := make([][]int, numCourses*numCourses)

	for i := 0; i < numCourses; i++ {
		graphTable[i] = make([]int, 0, numCourses)
	}

	for _, i := range prerequisites {
		before, after := i[0], i[1]
		graphTable[before] = append(graphTable[before], after)
	}

	checked := make([]bool, numCourses)
	visited := make([]bool, numCourses)
	result := make([]int, 0, numCourses)
	for i := 0; i < numCourses; i++ {
		if !checked[i] && hasCycleFindOrder(graphTable, visited, checked, &result, i) {
			return []int{}
		}
	}
	return result
}

//func findOrderTOPOSort(numCourses int, prerequisites [][]int) []int {
//	if numCourses == 0 {
//		return true
//	}
//	if len(prerequisites) == 0{
//		return true
//	}
//
//	graphTable := make([][]int, numCourses * numCourses)
//	inDegree := make([]int, numCourses)
//
//	for i:=0;i<numCourses;i++{
//		graphTable[i] = make([]int, 0, numCourses)
//	}
//
//	for _, i := range prerequisites{
//		after, before := i[0], i[1]
//		graphTable[before] = append(graphTable[before], after)
//		inDegree[after] += 1
//	}
//
//	Queue := make([]int, 0, numCourses)
//
//	for i:=0;i<numCourses;i++{
//		if inDegree[i] == 0{
//			Queue = append(Queue, i)
//		}
//	}
//
//	count := 0
//	for len(Queue) >0 {
//		v := Queue[0]
//		Queue = Queue[1:]
//		count += 1
//		nextNodeList := graphTable[v]
//		for _, i := range nextNodeList{
//			inDegree[i] -= 1
//			if inDegree[i] == 0{
//				Queue = append(Queue, i)
//			}
//		}
//
//	}
//	return count == numCourses
//}
