package RandomGen

import (
	"math"
	"math/rand"
)

func rand7() int {
	return int(rand.Intn(6)) + 1
}

func rand10() int {
	x := int(math.Pow(2, 32))
	for x > 40 {
		x = 7*(rand7()-1) + rand7()
	}
	return x%10 + 1
}
