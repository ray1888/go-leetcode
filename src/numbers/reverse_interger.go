package numbers

import "sync"

const (
	IntMax = 2 << 30
	IntMin = -(2 << 30)
)

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	y := 0
	for x != 0 {
		y = (x % 10) + (y * 10)
		x = x / 10
		if y >= IntMax || y <= IntMin {
			return 0
		}
	}
	a := sync.Map{}
	return y
}
