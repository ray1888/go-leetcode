package bit

func reverseBits(num uint32) uint32 {
	var result uint32
	var i uint32
	for ; i < 32; i++ {
		result = (result << 1) | (num & 1)
		num >>= 1
	}
	return result
}
