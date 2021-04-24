package main

// "divide and conquer" strategy that is introduced in "Hacker's Delight"
func hammingWeight(num uint32) int {
	num = (num & 0x55555555) + ((num >> 1) & 0x55555555)
	num = (num & 0x33333333) + ((num >> 2) & 0x33333333)
	num = (num + (num >> 4)) & 0x0f0f0f0f
	num = num + (num >> 8)
	num = num + (num >> 16)
	return int(num) & (1<<7 - 1)
}
