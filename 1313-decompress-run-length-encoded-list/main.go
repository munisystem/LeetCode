package main

func decompressRLElist(nums []int) []int {
	ans := []int{}
	for i := 0; i*2 < len(nums); i++ {
		for j := 0; j < nums[i*2]; j++ {
			ans = append(ans, nums[i*2+1])
		}
	}
	return ans
}
