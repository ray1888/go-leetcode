package numbers

import "math"

func isSquare(input int) bool {
	x := int(math.Sqrt(float64(input)))
	return x*x == input
}

func SumOfSqaureNumberByDiff(input int) bool {
	if input < 0 {
		return false
	}
	sqrtC := int(math.Sqrt(float64(input)))
	for i := 0; i <= sqrtC; i++ {
		if isSquare(input - i*i) {
			return true
		}
	}
	return false
}

func SumOfSquareNumberByMap(input int) bool {
	if input < 0 {
		return false
	}
	sqrtC := int(math.Sqrt(float64(input)))
	SquareMap := make(map[int]bool)
	for i := 0; i <= sqrtC; i++ {
		sqaure := i * i
		if _, ok := SquareMap[sqaure]; !ok {
			SquareMap[sqaure] = true
		}
		if _, ok := SquareMap[input-sqaure]; ok {
			return true
		}
	}
	return false
}

func SumOfSquareNumberTwoPointer(input int) bool {
	if input < 0 {
		return false
	}
	i := 0
	j := int(math.Sqrt(float64(input)))

	for i <= j {
		sumSqaure := i*i + j*j
		if sumSqaure == input {
			return true
		} else if sumSqaure < input {
			i += 1
		} else {
			j -= 1
		}
	}
	return false
}
