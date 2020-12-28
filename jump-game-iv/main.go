package main

func minJumps(arr []int) int {
	if len(arr) == 1 {
		return 0
	}
	dist := make([]int, len(arr), len(arr))
	for i := 0; i < len(dist); i++ {
		dist[i] = -1
	}
	queue := []int{}
	dist[0] = 0
	queue = append(queue, 0)
	values := map[int]int{}
	for len(queue) != 0 {
		i := queue[0]
		queue = queue[1:]
		if i == len(arr)-1 {
			break
		}
		if i > 0 {
			if dist[i-1] == -1 {
				dist[i-1] = dist[i] + 1
				queue = append(queue, i-1)
			}
		}
		if i < len(arr) {
			if dist[i+1] == -1 {
				dist[i+1] = dist[i] + 1
				queue = append(queue, i+1)
			}
		}
		if values[arr[i]] == 1 {
			continue
		}
		for j := 0; j < len(arr); j++ {
			if j == i || arr[j] != arr[i] {
				continue
			}
			if dist[j] == -1 {
				dist[j] = dist[i] + 1
				queue = append(queue, j)
			}
		}
		values[arr[i]] = 1
	}
	return dist[len(arr)-1]
}
