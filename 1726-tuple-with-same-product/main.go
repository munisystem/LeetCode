package main

func tupleSameProduct(nums []int) int {
	m := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			m[nums[i]*nums[j]]++
		}
	}
	r := 0
	for _, v := range m {
		if v == 1 {
			continue
		}
		r += 4 * (v * (v - 1))
	}
	return r
}
