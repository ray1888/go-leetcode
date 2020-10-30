package bit

func reverseBits(num uint32) uint32 {
	power, ret := uint32(31), uint32(0)
	for num != 0 {
		ret += (num & 1) << power
		num = num >> 1
		power--
	}
	return ret
}
