package numbers

import (
	"math"
)

func PowerOf3Brutal(input int) bool {
	if input == 0 {
		return false
	}
	for input%3 == 0 {
		input /= 3
	}
	return input == 1
}

var MaxPower = int(math.Log(2<<30) / math.Log(3))
var MaxNum = int(math.Pow(3, float64(MaxPower)))

func PowerOf3Math(input int) bool {
	return input > 0 && (MaxNum%input) == 0
}
