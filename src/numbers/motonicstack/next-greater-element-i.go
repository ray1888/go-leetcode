package montonicstack

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}
	m := make(map[int]int)
	var stack []int
	for _, v := range nums2 {
		for len(stack) != 0 && v > stack[len(stack)-1] {
			// 发现有更大的数字，给其下一个更大数字赋值
			m[stack[len(stack)-1]] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, v)
	}
	for k, v := range nums1 {
		if value, ok := m[v]; ok {
			nums1[k] = value
		} else {
			nums1[k] = -1
		}
	}
	return nums1
}
