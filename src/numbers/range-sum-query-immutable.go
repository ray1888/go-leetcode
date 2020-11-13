package numbers

type NumArray struct {
	PrefixSum []int
}

func constructor(nums []int) NumArray {
	n := NumArray{}
	n.PrefixSum = make([]int, len(nums)+1)
	for i := 1; i < len(n.PrefixSum); i++ {
		n.PrefixSum[i] = n.PrefixSum[i-1] + nums[i-1]
	}
	return n
}

func (this *NumArray) sumRange(i int, j int) int {
	return this.PrefixSum[j+1] - this.PrefixSum[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
