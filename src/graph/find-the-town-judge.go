package graph

func findJudge(N int, trust [][]int) int {

	inDegree := make([]int, N+1)
	outDegree := make([]int, N+1)

	for i := 0; i < len(trust); i++ {
		inDegree[trust[i][1]]++
		outDegree[trust[i][0]]++
	}

	for i := 1; i <= N; i++ {
		if inDegree[i] == N-1 && outDegree[i] == 0 {
			return i
		}
	}
	return -1
}
