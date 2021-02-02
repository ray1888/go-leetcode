package DynamicProgramming

/*
	如果没有 00 的话，直接考虑维护一个前缀积 pre[i]pre[i] 表示前 ii 个数的乘积即可，答案就是 \frac{pre[n]}{pre[n-k]}
pre[n−k]
pre[n]
​
 ，其中 nn 表示当前 prepre 数组的长度。那么如何处理 00 呢？可以注意到如果出现 00 的话，那么 00 之前的数对答案都是没有用的了，所以我们可以遇到 00 的时候直接清空 prepre 数组，那么询问的时候我们要求的是末尾 kk 个数的乘积，如果这时候我们 prepre 数组的长度小于 kk，说明末尾 kk 个数里肯定有 00，直接输出 00 即可，否则输出 \frac{pre[n]}{pre[n-k]}
pre[n−k]
pre[n]
​


作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/product-of-the-last-k-numbers/solution/zui-hou-k-ge-shu-de-cheng-ji-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

type ProductOfNumbers struct {
	prefixProductArray []int
	length             int
}

func ConstructorProductPrefix() ProductOfNumbers {
	p := ProductOfNumbers{}
	p.prefixProductArray = make([]int, 40001)
	p.prefixProductArray[0] = 1
	p.length = 0
	return p
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.length = 0
	} else {
		this.length++
		this.prefixProductArray[this.length] = num
		this.prefixProductArray[this.length] *= this.prefixProductArray[this.length-1]
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	if this.length < k {
		return 0
	}
	return this.prefixProductArray[this.length] / this.prefixProductArray[this.length-k]
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
