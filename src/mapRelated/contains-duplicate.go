package mapRelated

func containsDuplicate(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	m := make(map[int]bool)
	for _, num := range nums {
		if _, ok := m[num]; !ok {
			m[num] = true
		} else {
			return true
		}
	}
	return false
}
