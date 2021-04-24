package main

func sortColors(nums []int) {
	sort(nums, 0, len(nums)-1)
}

func sort(nums []int, left, right int) {
	if (right - left) < 1 {
		return
	}
	pi := (left + right) / 2
	p := nums[pi]
	swap(nums, pi, right)

	i := left
	for j := left; j <= right; j++ {
		if nums[j] < p {
			swap(nums, i, j)
			i++
		}
	}
	swap(nums, i, right)
	sort(nums, left, i)
	sort(nums, i+1, right)
}

func swap(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}
