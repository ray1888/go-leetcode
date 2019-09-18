package DynamicProgramming

func minCostClimbingStairs(cost []int) int {
	if len(cost) == 0 {
		return 0
	}

	first := cost[0]
	second := cost[1]
	cur := 0
	for start := 2; start < len(cost); start++ {
		cur = min(first, second) + cost[start]
		first = second
		second = cur
	}
	return min(first, second)
}
