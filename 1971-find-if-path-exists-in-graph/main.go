package main

func validPath(n int, edges [][]int, start int, end int) bool {
	queue := []int{}
	queue = append(queue, start)
	visited := make([]bool, n)
	for len(queue) != 0 {
		p := queue[0]
		queue = queue[1:]
		if p == end {
			return true
		}
		if visited[p] {
			continue
		}
		visited[p] = true
		for _, edge := range edges {
			if edge[0] == p {
				queue = append(queue, edge[1])
			} else if edge[1] == p {
				queue = append(queue, edge[0])
			}
		}
	}
	return false
}
