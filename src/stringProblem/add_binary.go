package stringProblem

import "strconv"

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		//始终让len(a)>len(b)
		a, b = b, a
	}
	var str = ""
	//需要进位 则变为 1
	var flag = 0
	var lenAB = len(a) - len(b)
	//补零操作
	if len(a) != len(b) {
		for i := 0; i < lenAB; i++ {
			b = "0" + b
		}
	}

	for i := len(a) - 1; i >= 0; i-- {
		aNum := a[i] - '0'
		bNum := b[i] - '0'

		sum := int(aNum+bNum) + flag
		if sum == 2 {
			str = "0" + str
			flag = 1
		} else if sum == 3 {
			str = "1" + str
			flag = 1
		} else {
			str = strconv.Itoa(sum) + str
			flag = 0
		}
	}
	//确保最后一位需要进位
	if flag == 1 {
		str = "1" + str
	}
	return str
}
