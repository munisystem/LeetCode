package main

func jump(nums []int) int {
	d := make([]int, len(nums), len(nums))
	for i := 0; i < len(nums); i++ {
		d[i] = -1
	}
	queue := []int{}
	queue = append(queue, 0)
	d[0] = 0
	for len(queue) != 0 {
		i := queue[0]
		queue = queue[1:]
		for j := 1; j <= nums[i]; j++ {
			if i+j < len(nums) && d[i+j] == -1 {
				d[i+j] = d[i] + 1
				queue = append(queue, i+j)
			}
			if i+j == len(nums)-1 {
				break
			}
		}
	}
	return d[len(d)-1]
}
