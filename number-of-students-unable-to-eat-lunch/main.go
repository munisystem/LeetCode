package main

func countStudents(students []int, sandwiches []int) int {
	m := make(map[int]int, 2)
	for i := 0; i < len(students); i++ {
		m[students[i]]++
	}
	ans := 0
	for i := 0; i < len(sandwiches); i++ {
		if m[sandwiches[i]] > 0 {
			m[sandwiches[i]]--
			ans++
		} else {
			break
		}
	}
	return len(sandwiches) - ans
}
