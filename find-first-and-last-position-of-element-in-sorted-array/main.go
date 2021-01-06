package main

func searchRange(nums []int, target int) []int {
	start, end, m := 0, len(nums)-1, 0
	for start <= end {
		m = (start + end) / 2
		if nums[m] == target {
			i, j := m, m
			for ; i > 0; i-- {
				if nums[i-1] != target {
					break
				}
			}
			for ; j < len(nums)-1; j++ {
				if nums[j+1] != target {
					break
				}
			}
			return []int{i, j}
		}
		if nums[m] > target {
			end = m - 1
		} else {
			start = m + 1
		}
	}
	return []int{-1, -1}
}
