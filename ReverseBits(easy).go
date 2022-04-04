package leetcode

func reverseBits(num uint32) uint32 {
	var res uint32
	i := 0
	for i < 32 {
		res = res<<1 | num&1
		num = num >> 1
		i++
	}
	return res
}
