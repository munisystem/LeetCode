package main

func uniqueOccurrences(arr []int) bool {
	occurrences := map[int]int{}
	for i := 0; i < len(arr); i++ {
		occurrences[arr[i]]++
	}
	m := map[int]interface{}{}
	for _, v := range occurrences {
		if _, ok := m[v]; ok {
			return false
		}
		m[v] = struct{}{}
	}
	return true
}
