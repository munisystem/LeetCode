package main

func findGCD(nums []int) int {
	min, max := nums[0], nums[0]
	for _, num := range nums {
		if min > num {
			min = num
		}
		if max < num {
			max = num
		}
	}

	for i := min; i >= 2; i-- {
		if min%i == 0 && max%i == 0 {
			return i
		}
	}
	return 1
}
