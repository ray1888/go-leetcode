package mapRelated

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}
	record := make(map[int]int)
	result := make([]int, 0, 10)
	for _, val := range nums1 {
		if _, ok := record[val]; !ok {
			record[val] = 1
		} else {
			record[val] += 1
		}
	}
	for _, val := range nums2 {
		if _, ok := record[val]; !ok {
			continue
		}
		if record[val] > 0 {
			result = append(result, val)
			record[val] -= 1
		}
	}
	return result
}

func intersectSort(nums1, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}
	sort.Ints(nums1)
	sort.Ints(nums2)

	i := 0
	j := 0
	len1 := len(nums1)
	len2 := len(nums2)
	result := make([]int, 0, 10)
	for i < len1 && j < len2 {
		if nums1[i] == nums2[j] {
			result = append(result, nums1[i])
			i += 1
			j += 1
		} else {
			if nums1[i] > nums2[j] {
				j += 1
			} else {
				i += 1
			}
		}
	}
	return result
}
