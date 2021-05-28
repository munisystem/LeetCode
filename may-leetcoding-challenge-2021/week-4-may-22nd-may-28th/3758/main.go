package main

import "fmt"

func main() {
	fmt.Println(maximumUniqueSubarray([]int{4, 2, 4, 5, 6}))
}

func maximumUniqueSubarray(nums []int) int {
	var sum, ans int
	var i, j int
	m := map[int]int{}
	for i < len(nums) && j < len(nums) {
		sum += nums[j]
		m[nums[j]]++
		if m[nums[j]] <= 1 {
			if ans < sum {
				ans = sum
			}
			j++
			continue
		}
		for i < j {
			sum -= nums[i]
			m[nums[i]]--
			cnt := m[nums[i]]
			i++
			if cnt == 1 {
				break
			}
		}
		j++
	}
	return ans
}
