package DynamicProgramming

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	n1 := len(s1)
	n2 := len(s2)
	d := make([][]bool, n1+1)
	for index := range d {
		d[index] = make([]bool, n2+1)
	}
	d[0][0] = true

	for i := 1; i <= n1; i++ {
		d[i][0] = d[i-1][0] && s1[i-1] == s3[i-1]
	}
	for j := 1; j <= n2; j++ {
		d[0][j] = d[0][j-1] && s2[j-1] == s3[j-1]
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			k := i + j
			d[i][j] = (d[i-1][j] && s1[i-1] == s3[k-1]) || (d[i][j-1] && s2[j-1] == s3[k-1])
		}
	}
	return d[n1][n2]
}
