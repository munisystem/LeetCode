package main

func numJewelsInStones(jewels string, stones string) int {
	m := make(map[byte]interface{}, len(jewels))
	for i := 0; i < len(jewels); i++ {
		m[jewels[i]] = struct{}{}
	}
	ans := 0
	for i := 0; i < len(stones); i++ {
		if _, ok := m[stones[i]]; ok {
			ans++
		}
	}
	return ans
}
