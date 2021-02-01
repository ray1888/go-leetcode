package DynamicProgramming

type NumArray struct {
	PrefixSum []int
}

func Constructor(nums []int) NumArray {
	n := NumArray{}
	n.PrefixSum = make([]int, len(nums)+1)
	for i := 1; i < len(n.PrefixSum); i++ {
		n.PrefixSum[i] = n.PrefixSum[i-1] + nums[i-1]
	}
	// fmt.Println(n.PrefixSum)
	return n
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.PrefixSum[j+1] - this.PrefixSum[i]
}
