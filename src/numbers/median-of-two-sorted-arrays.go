package numbers

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	if (total & 1) == 1 {
		return float64(findKthSmallestInSortedArrays(nums1, nums2, total/2+1))
	} else {
		a := findKthSmallestInSortedArrays(nums1, nums2, total/2)
		b := findKthSmallestInSortedArrays(nums1, nums2, total/2+1)
		return float64(a+b) / 2
	}
}

func findKthSmallestInSortedArrays(nums1, nums2 []int, k int) int {
	length1 := len(nums1)
	length2 := len(nums2)
	base1 := 0
	base2 := 0
	for true {
		if length1 == 0 {
			return nums2[base2+k-1]
		}
		if length2 == 0 {
			return nums1[base1+k-1]
		}
		if k == 1 {
			return min(nums1[base1], nums2[base2])
		}

		i := min(k/2, length1)
		j := min(k-i, length2)
		a := nums1[base1+i-1]
		b := nums2[base2+j-1]
		if i+j == k && a == b {
			return a
		}

		if a >= b {
			base2 += j
			length2 -= j
			k -= j
		}

		if b >= a {
			base1 += i
			length1 -= i
			k -= i
		}
	}
	return -1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
