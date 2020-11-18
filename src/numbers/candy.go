package numbers

func candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}
	n := len(ratings)
	d := make([]int, n)
	d[0] = 1
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			d[i] = d[i-1] + 1
		} else {
			d[i] = 1
		}
	}
	sum := d[n-1]
	for j := n - 2; j >= 0; j-- {
		if ratings[j] > ratings[j+1] {
			d[j] = max(d[j+1]+1, d[j])
		}
		sum += d[j]
	}
	return sum
}
