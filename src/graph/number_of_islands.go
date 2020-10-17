package graph

func dfs(graph [][]byte, visited [][]bool, i, j int) {
	row := len(graph) - 1
	col := len(graph[0]) - 1
	if i < 0 || j < 0 || i > row || j > col || graph[i][j] == '0' || visited[i][j] {
		return
	}
	visited[i][j] = true
	dfs(graph, visited, i+1, j)
	dfs(graph, visited, i-1, j)
	dfs(graph, visited, i, j+1)
	dfs(graph, visited, i, j-1)
	return
}

func numIslands(grid [][]byte) int {
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
				dfs(grid, visits, i, j)
			}
		}
	}
	return number
}
