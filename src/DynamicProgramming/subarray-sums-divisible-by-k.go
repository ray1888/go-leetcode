package DynamicProgramming

func subarraysDivByK(A []int, K int) int {
	if len(A) == 0 {
		return 0
	}
	m := make(map[int]int)
	sum := 0
	res := 0
	m[0] = 1
	for i := 0; i < len(A); i++ {
		sum += A[i]
		modulus := (sum%K + K) % K
		res += m[modulus]
		m[modulus]++
	}
	return res
}

func subarraysDivByKSingleCount(A []int, K int) int {
	record := map[int]int{0: 1}
	sum, ans := 0, 0
	for _, elem := range A {
		sum += elem
		modulus := (sum%K + K) % K
		record[modulus]++
	}
	for _, cx := range record {
		ans += cx * (cx - 1) / 2
	}
	return ans
}
