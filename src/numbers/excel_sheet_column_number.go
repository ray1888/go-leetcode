package numbers

func titleToNumberRight2Left(s string) int {
	num := 0
	base := 1
	for i := len(s) - 1; i >= 0; i-- {
		num += int(uint8(s[i])-'A'+1) * base
		base *= 26
	}
	return num
}

func titleToNumberLeft2Right(s string) int {
	num := 0
	for i := 0; i < len(s); i++ {
		num = num*26 + int(uint8(s[i])-'A'+1)
	}
	return num
}
