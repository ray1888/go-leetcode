package numbers

import "go-leetcode/src/utils"

func findKthNumber(n int, k int) int {
	getCount := func(prefix int) int {
		cur, next := prefix, prefix+1
		count := 0
		for cur <= n {
			// 下一个前缀的起点减去当前前缀的起点就是 prefix 开始到 n 之间的本层节点数量
			count += utils.Min(next, n+1) - cur
			cur *= 10
			next *= 10
		}
		return count
	}
	preNum := 1
	parent := 1
	for preNum < k {
		count := getCount(parent)
		if count+preNum > k {
			// 此时说明满足条件的元素在 prefix 字典树下，转到 prefix 子树下寻找
			parent *= 10
			// prefix元素的序号是要统计的
			preNum++
		} else {
			// 此时说明在当前这一层的某棵子树下，我们给prefix++，表示寻找下一棵同级子树
			parent++
			// 并且统计元素数量
			preNum += count
		}
	}
	return parent
}
