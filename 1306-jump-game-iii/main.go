package main

func canReach(arr []int, start int) bool {
	d := make([]bool, len(arr), len(arr))
	zeroi := []int{}
	for i := 0; i < len(d); i++ {
		d[i] = false
		if arr[i] == 0 {
			zeroi = append(zeroi, i)
		}
	}
	d[start] = true
	queue := []int{}
	queue = append(queue, start)

	for len(queue) != 0 {
		i := queue[0]
		queue = queue[1:]
		j := arr[i]
		if i-j >= 0 && !d[i-j] {
			d[i-j] = true
			queue = append(queue, i-j)
		}
		if i+j < len(d) && !d[i+j] {
			d[i+j] = true
			queue = append(queue, i+j)
		}
	}
	for _, i := range zeroi {
		if d[i] {
			return true
		}
	}
	return false
}
