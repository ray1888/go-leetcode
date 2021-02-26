package graph

func dfsIslandII(graph [][]byte, visited [][]bool, i, j int) {
	row := len(graph) - 1
	col := len(graph[0]) - 1
	if i < 0 || j < 0 || i > row || j > col || graph[i][j] == '0' || visited[i][j] {
		return
	}
	visited[i][j] = true
	dfsIsland(graph, visited, i+1, j)
	dfsIsland(graph, visited, i-1, j)
	dfsIsland(graph, visited, i, j+1)
	dfsIsland(graph, visited, i, j-1)
	return
}

func numIslandsDFS(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	visits := make([][]bool, len(grid))
	for index := range visits {
		visits[index] = make([]bool, len(grid[0]))
	}

	row := len(grid)
	col := len(grid[0])
	var number int
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' && !visits[i][j] {
				number++
				dfsIsland(grid, visits, i, j)
			}
		}
	}
	return number
}

func numberIslandsIIONSquare(m int, n int, positions [][]int) []int {
	grid := make([][]byte, m)
	for i := 0; i < n; i++ {
		grid[i] = make([]byte, n)
	}
	ans := make([]int, 0)
	for _, item := range positions {
		grid[item[0]][item[1]] = '1'
		ans = append(ans, numIslandsDFS(grid))
	}
	return ans

}

func numberIslandUnion(m, n int, posititions [][]int) []int {
	parent := make([]int, m*n)
	for i := 0; i < len(parent); i++ {
		parent[i] = -1
	}
	count := 0
	ans := make([]int, 0)
	for _, positition := range posititions {
		i := positition[0]*n + positition[1]
		if parent[i] != -1 {
			ans = append(ans, count)
			continue
		}
		parent[i] = i
		count++
		if positition[0] > 0 && parent[i-n] != -1 {
			count = union(i, i-n, parent, count)
		}
		if positition[0] < m-1 && parent[i+n] != -1 {
			count = union(i, i+n, parent, count)
		}
		if positition[1] > 0 && parent[i-1] != -1 {
			count = union(i, i-1, parent, count)
		}
		if positition[1] < n-1 && parent[i+1] != -1 {
			count = union(i, i+1, parent, count)
		}
		ans = append(ans, count)

	}
	return ans
}

func union(i, j int, parent []int, count int) int {
	parent1 := find(i, parent)
	parent2 := find(j, parent)
	if parent1 != parent2 {
		count--
		parent[parent2] = parent1
	}
	return count
}

func find(i int, parent []int) int {
	if parent[i] != i {
		parent[i] = find(parent[i], parent)
	}
	return parent[i]
}
