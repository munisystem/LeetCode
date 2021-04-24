package main

func numWaterBottles(numBottles int, numExchange int) int {
	ans, emptyBottles := 0, 0
	for numBottles > 0 {
		ans += numBottles
		emptyBottles += numBottles
		numBottles, emptyBottles = emptyBottles/numExchange, emptyBottles%numExchange
	}
	return ans
}
