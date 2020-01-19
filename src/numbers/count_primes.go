package numbers

import "math"

func CountPrime(n int) int {
	if n <= 2 {
		return 0
	}
	primes := make([]bool, n)
	for i := 0; i < n; i++ {
		primes[i] = true
	}
	primes[0] = false
	primes[1] = false
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrtN; i++ {
		if primes[i] {
			for j := i * i; j < n; j += i {
				primes[j] = false
			}
		}
	}
	count := 0
	for _, val := range primes {
		if val == true {
			count += 1
		}
	}
	return count
}

func CountPrimeFast(n int) int {
	if n <= 2 {
		return 0
	}
	composites := make([]bool, n)
	count := 1
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 3; i < n; i += 2 {
		if composites[i] {
			continue
		}
		count += 1
		if i > sqrtN {
			continue
		}
		for j := i * i; j < n; j += 2 * i {
			composites[j] = true
		}
	}
	return count
}
