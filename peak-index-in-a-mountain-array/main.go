package main

func peakIndexInMountainArray(arr []int) int {
	start, end := 0, len(arr)-1
	for start < end {
		mid := (start + end) / 2
		if arr[mid] < arr[mid+1] {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}
