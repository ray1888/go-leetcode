package DynamicProgramming

type NumMatrix struct {
	PrefixSum [][]int // 保存二维矩阵的所有前缀和
}

func Constructor2(matrix [][]int) NumMatrix {
	// 首先处理边界情况
	if matrix == nil || len(matrix) == 0 || matrix[0] == nil || len(matrix[0]) == 0 {
		prefixSum := make([][]int, 1) // 把前缀和定义为1x1的二维矩阵
		for i := range prefixSum {
			prefixSum[i] = make([]int, 1)
		}
		return NumMatrix{PrefixSum: prefixSum}
	}
	// 否则把矩阵的两个维度取出,分别赋值给m和n,方便使用
	m, n := len(matrix), len(matrix[0])
	// 并把前缀和数组定义为 m+1 * n+1 的二维数组
	prefixSum := make([][]int, m+1)
	for i := range prefixSum {
		prefixSum[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			prefixSum[i][j] = prefixSum[i][j-1] + prefixSum[i-1][j] - prefixSum[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	return NumMatrix{PrefixSum: prefixSum}
}

// 子矩阵求和
func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	// 坐标row1,col1到坐标row2,col2子矩阵内的元素和
	return this.PrefixSum[row2+1][col2+1] - this.PrefixSum[row2+1][col1] - this.PrefixSum[row1][col2+1] + this.PrefixSum[row1][col1]
}
