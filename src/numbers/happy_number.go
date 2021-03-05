package numbers

func transform(input int) int {
	sum := 0
	for input != 0 {
		a := input % 10
		sum += a * a
		input /= 10
	}
	return sum
}

func HappyNubmerTwoPointer(x int) bool {
	if x <= 0 {
		return false
	}
	fast := x
	slow := x
	for true {
		fast = transform(transform(fast))
		slow = transform(slow)
		if fast == 1 {
			return true
		}
		if fast == slow {
			return false
		}
	}
	return true
}

func isHappy(n int) bool {
	m := map[int]bool{}
	for ; n != 1 && !m[n]; n, m[n] = step(n), true {
	}
	return n == 1
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}
