package main

func relativeSortArray(arr1 []int, arr2 []int) []int {
	nums := make([]int, 1001, 1001)
	for _, v := range arr1 {
		nums[v]++
	}
	ans := make([]int, 0, len(arr1))
	for _, v := range arr2 {
		for nums[v] > 0 {
			ans = append(ans, v)
			nums[v]--
		}
	}
	for k, v := range nums {
		for i := 0; i < v; i++ {
			ans = append(ans, k)
		}
	}
	return ans
}
