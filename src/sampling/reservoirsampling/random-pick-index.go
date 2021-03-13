package reservoirsampling

import (
	"math/rand"
	"time"
)

type SolutionNums struct {
	r    *rand.Rand
	nums []int
}

func ConstructorNums(nums []int) SolutionNums {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	solution := SolutionNums{
		r:    r,
		nums: nums,
	}
	return solution
}

func (this *SolutionNums) Pick(target int) int {
	k := 1
	var res int
	for i := 0; i < len(this.nums); i++ {
		if this.nums[i] == target {
			if this.r.Intn(k) == 0 {
				res = i
			}
			k++
		}

	}
	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
