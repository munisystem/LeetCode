package main

func findCenter(edges [][]int) int {
	m := map[int]int{}
	for i := 0; i < len(edges); i++ {
		m[edges[i][0]]++
		m[edges[i][1]]++
	}
	var ans int
	for i, v := range m {
		if v == len(edges) {
			ans = i
			break
		}
	}
	return ans
}
