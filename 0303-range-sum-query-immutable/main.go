package main

type NumArray struct {
	cumsum []int
}

func Constructor(nums []int) NumArray {
	cumsum := make([]int, len(nums)+1, len(nums)+1)
	cumsum[0] = 0
	for i := 0; i < len(nums); i++ {
		cumsum[i+1] = cumsum[i] + nums[i]
	}
	return NumArray{cumsum: cumsum}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.cumsum[right+1] - this.cumsum[left]
}
