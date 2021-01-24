package numbers

func maxPointsBrutulForce(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	maxNum := 1
	for i := 0; i < len(points); i++ {
		var ax, ay int
		ax = points[i][0]
		ay = points[i][1]
		samp := 0
		for j := i + 1; j < len(points); j++ {
			var bx, by int
			bx = points[j][0]
			by = points[j][1]
			if ax == bx && ay == by {
				samp++
				continue
			}
			cur := 2
			for k := j + 1; k < len(points); k++ {
				var cx, cy int
				cx = points[k][0]
				cy = points[k][1]
				if (by-ay)*(cx-ax)-(bx-ax)*(cy-ay) == 0 {
					cur++
				}
			}
			maxNum = max(cur+samp, maxNum)
		}
		maxNum = max(maxNum, samp+1)
	}
	return maxNum
}

func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func maxPoints(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	maxNum := 1
	for i := 0; i < len(points); i++ {
		var ax, ay int64
		ax = int64(points[i][0])
		ay = int64(points[i][1])
		samp := 0
		lineMax := 0
		result := make(map[int64]int)
		for j := i + 1; j < len(points); j++ {
			var bx, by int64
			bx = int64(points[j][0])
			by = int64(points[j][1])
			if ax == bx && ay == by {
				samp++
				continue
			}
			x := bx - ax
			y := by - ay
			g := gcd(y, x)
			x /= g
			y /= g
			key := (x << 32) + y
			if _, ok := result[key]; !ok {
				result[key] = 0
			}
			count := result[key] + 1
			result[key] = count
			lineMax = max(count, lineMax)
		}
		maxNum = max(maxNum, lineMax+samp+1)
	}
	return maxNum
}
