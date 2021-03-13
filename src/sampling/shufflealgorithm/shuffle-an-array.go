package shufflealgorithm

import (
	"math/rand"
	"time"
)

type Solution struct {
	origin []int
	new    []int
	r      *rand.Rand
}

func Constructor(nums []int) Solution {
	s := Solution{}
	s.origin = make([]int, len(nums))
	s.new = make([]int, len(nums))
	s.r = rand.New(rand.NewSource(time.Now().UnixNano()))
	copy(s.origin, nums)
	copy(s.new, nums)
	return s
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	copy(this.new, this.origin)
	return this.new
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	n := len(this.new)
	for i := n - 1; i >= 0; i-- {
		randNum := this.r.Intn(i + 1)                                   // math.rand中的Intn(i+1)返回[0, i]范围的整数，每次数组在下标index为[0, i]范围内随机找一个下标对应的元素与当前位置i处的元素进行交换
		this.new[i], this.new[randNum] = this.new[randNum], this.new[i] // 对应位置元素交换，也可以使用如下代码
	}
	return this.new
}
