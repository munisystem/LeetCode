package main

func canBeEqual(target []int, arr []int) bool {
	m := map[int]int{}
	for i := 0; i < len(target); i++ {
		m[target[i]]++
		m[arr[i]]++
	}
	for _, v := range m {
		if v%2 != 0 {
			return false
		}
	}
	return true
}
