package main

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		diff := target - (numbers[i] + numbers[j])
		if diff == 0 {
			break
		}
		if diff < 0 {
			j--
		} else {
			i++
		}
	}
	return []int{i + 1, j + 1}
}
