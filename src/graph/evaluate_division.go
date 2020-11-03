package graph

func dfsED(start string, end string, result float64, visited map[string]bool, graph map[string]map[string]float64) float64 {
	nexts, ok := graph[start]
	if !ok {
		return -1
	}
	if val, ok := nexts[end]; ok {
		return val * result
	}
	// fmt.Println(start)
	for key, item := range nexts {
		// fmt.Println(key)
		if val, ok := visited[key]; ok && val {
			continue
		}
		tmp := result * item
		visited[key] = true
		ret := dfsED(key, end, tmp, visited, graph)
		if ret != -1 {
			return ret
		}
	}
	return -1
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string]map[string]float64)
	for i := 0; i < len(equations); i++ {
		equation := equations[i]
		if _, ok := graph[equation[0]]; !ok {
			graph[equation[0]] = make(map[string]float64)
		}
		graph[equation[0]][equation[1]] = values[i]
		if _, ok := graph[equation[1]]; !ok {
			graph[equation[1]] = make(map[string]float64)
		}
		graph[equation[1]][equation[0]] = 1 / values[i]
	}
	result := make([]float64, 0)

	for i := 0; i < len(queries); i++ {
		query := queries[i]
		start := query[0]
		end := query[1]
		if _, ok := graph[end]; !ok {
			result = append(result, -1)
			continue
		}
		if _, ok := graph[start]; !ok {
			result = append(result, -1)
			continue
		}
		if start == end {
			result = append(result, 1)
			continue
		}
		visited := make(map[string]bool)
		ret := dfsED(start, end, 1.0, visited, graph)
		result = append(result, ret)
	}
	return result
}
