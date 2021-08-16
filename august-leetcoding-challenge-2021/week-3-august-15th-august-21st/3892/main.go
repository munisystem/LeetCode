package main

type NumArray struct {
	array []int
}

func Constructor(nums []int) NumArray {
	array := make([]int, len(nums)+1, len(nums)+1)
	array[0] = 0
	for i := 0; i < len(nums); i++ {
		array[i+1] = array[i] + nums[i]
	}
	return NumArray{array: array}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.array[right+1] - this.array[left]
}
