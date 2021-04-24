package main

func search(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start < end {
		m := (start + end) / 2
		p := nums[m]
		e := nums[end]
		if p == target {
			return m
		}
		if e == target {
			return end
		}
		if p < e {
			switch {
			case target > p && target < e:
				start = m + 1
			default:
				end = m
			}
		} else {
			switch {
			case target < p && target > e:
				end = m
			default:
				start = m + 1
			}
		}
	}
	if nums[end] == target {
		return end
	}
	return -1
}
