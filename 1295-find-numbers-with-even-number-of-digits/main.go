package main

func findNumbers(nums []int) int {
	evens := 0
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		digits := 0
		for ; num > 0; num = num / 10 {
			digits = digits + 1
		}
		if digits%2 == 0 {
			evens = evens + 1
		}
	}
	return evens
}
