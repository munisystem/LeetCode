package main

func reverseBits(num uint32) uint32 {
	var ans uint32
	for i := 0; i < 32; i++ {
		ans += num >> i << 31 >> i
	}
	return ans
}
