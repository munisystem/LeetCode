package main

func canJump(nums []int) bool {
	dict := make([]bool, len(nums), len(nums))
	for i := 0; i < len(dict); i++ {
		dict[i] = false
	}
	dict[0] = true
	queue := []int{}
	queue = append(queue, 0)
	for len(queue) != 0 {
		i := queue[0]
		queue = queue[1:]
		for j := 1; j <= nums[i]; j++ {
			if i+j < len(dict) && !dict[i+j] {
				dict[i+j] = true
				queue = append(queue, i+j)
			}
			if i+j == len(nums)-1 {
				break
			}
		}
	}
	return dict[len(dict)-1]
}
