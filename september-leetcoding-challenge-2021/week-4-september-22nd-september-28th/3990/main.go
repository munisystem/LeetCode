package main

func sortArrayByParityII(nums []int) []int {
	even, odd := 0, 1
	for even < len(nums) && odd < len(nums) {
		for even < len(nums) && nums[even]%2 == 0 {
			even += 2
		}
		for odd < len(nums) && nums[odd]%2 != 0 {
			odd += 2
		}
		if even < len(nums) && odd < len(nums) {
			nums[even], nums[odd] = nums[odd], nums[even]
		}
	}
	return nums
}
