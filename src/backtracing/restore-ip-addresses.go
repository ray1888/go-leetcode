package backtracing

import "strconv"

var (
	segment    []int
	result     []string
	SegmentNum = 4
)

func restoreIpAddresses(s string) []string {
	if len(s) == 0 {
		return []string{}
	}
	result = make([]string, 0)
	segment = make([]int, SegmentNum)
	dfsIp(s, 0, 0)
	return result
}

/*
	此题的重点在于回溯把segment放入对应的位置上面
	假设输入为010010

	实际上的过程是
	0([0,0,0,0])-> 1 [0,1,0,0] -> 0 [0,1,0,0] -> 0 [0,1,0,0]
	然后到达SegmentID =4
	然后回溯
	0([0,0,0,0])-> 1 [0,1,0,0] -> 10 [0, 10, 0,0 ]
	如此类推
*/
func dfsIp(s string, segmentID int, segStartIndex int) {
	if segmentID == 4 {
		ipaddr := ""
		if segStartIndex == len(s) {
			for i := 0; i < SegmentNum; i++ {
				ip := strconv.Itoa(segment[i])
				ipaddr += ip
				if i != SegmentNum-1 {
					ipaddr += "."
				}
			}
			result = append(result, ipaddr)
		}
		return
	}
	if segStartIndex == len(s) {
		return
	}

	if s[segStartIndex] == '0' {
		segment[segmentID] = 0
		dfsIp(s, segmentID+1, segStartIndex+1)
	}

	addr := 0
	for segEnd := segStartIndex; segEnd < len(s); segEnd++ {
		addr = addr*10 + int(s[segEnd]-'0')
		if addr > 0 && addr <= 255 {
			segment[segmentID] = addr
			dfsIp(s, segmentID+1, segEnd+1)
		} else {
			break
		}
	}
}
