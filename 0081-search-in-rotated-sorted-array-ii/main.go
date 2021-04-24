package main

func search(nums []int, target int) bool {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return true
		}
		if nums[mid] == nums[end] {
			end--
			continue
		}

		if nums[mid] < nums[end] {
			if target > nums[mid] && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		} else {
			if target < nums[mid] && target > nums[end] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}
	return false
}
