package main

func kWeakestRows(mat [][]int, k int) []int {
	ans := make([]int, 0, k)
	m := map[int]interface{}{}
	for i := 0; i < len(mat[0]) && len(ans) < k; i++ {
		for j := 0; j < len(mat); j++ {
			if _, ok := m[j]; ok {
				continue
			}
			if mat[j][i] == 0 {
				ans = append(ans, j)
				m[j] = struct{}{}
				if len(ans) == k {
					break
				}
			}
		}
	}
	for i := 0; i < len(mat) && len(ans) < k; i++ {
		if _, ok := m[i]; ok {
			continue
		}
		ans = append(ans, i)
		m[i] = struct{}{}
	}
	return ans
}
