package maxnumber

import "strconv"

func maximumSwap(num int) int {

	str := []byte(strconv.Itoa(num))
	indexes := make([]int, 10)
	//记录每个元素的索引值，如果又重复，记录最右边的
	for i, b := range str {
		indexes[b-'0'] = i
	}
	//从左到右遍历str，然后每个元素和9-0对比，
	//例如：如果9存在，并且9的索引比当前大（在右边），那么换位置，返回结果
	for i, b := range str {
		for d := 9; d > int(b-'0'); d-- {
			if indexes[d] > i {
				str[i], str[indexes[d]] = str[indexes[d]], str[i]
				res, _ := strconv.Atoi(string(str))
				return res
			}
		}
	}
	return num
}
