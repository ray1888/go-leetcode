package numbers

import "go-leetcode/src/utils"

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	// 取 cg 的小值 减 ae 的大值 得到 边长
	w := utils.Min(C, G) - utils.Max(A, E)
	// 同理
	// 取 dh 的小值 减 bf 的大值 得到 高度
	h := utils.Min(D, H) - utils.Max(B, F)
	// 面值等于 s1+s2-重叠部分
	// 这里做下优化 避免溢出 写成 s1-重叠部分+s2
	// 上面算出来的w和h 有可能出现负数
	// 出现负数那就是没有重叠的部分了 置为0 去乘面积就好了 ， 那就等于s1+s2
	return (C-A)*(D-B) - utils.Max(w, 0)*utils.Max(h, 0) + (G-E)*(H-F)
}
