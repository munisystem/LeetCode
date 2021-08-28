package main

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	candies := map[int]struct{}{}
	var alice, bob int
	for _, v := range aliceSizes {
		candies[v] = struct{}{}
		alice += v
	}
	for _, v := range bobSizes {
		bob += v
	}
	ans := make([]int, 2, 2)
	total := (alice + bob) / 2
	for _, b := range bobSizes {
		if _, ok := candies[total-(bob-b)]; ok {
			ans[0], ans[1] = total-(bob-b), b
			break
		}
	}
	return ans
}
