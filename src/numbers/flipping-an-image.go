package numbers

func flipAndInvertImage(A [][]int) [][]int {
	for _, row := range A {
		i := 0
		j := len(row) - 1
		for i < j {
			if row[i] == row[j] {
				row[i] ^= 1
				row[j] ^= 1
			}
			i++
			j--
		}
		if i == j {
			row[i] ^= 1
		}
	}
	return A
}
