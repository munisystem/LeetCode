package main

func canVisitAllRooms(rooms [][]int) bool {
	queue := []int{}
	queue = append(queue, 0)
	m := make(map[int]interface{}, len(rooms))
	m[0] = struct{}{}
	for len(queue) != 0 {
		room := queue[0]
		queue = queue[1:]
		for i := 0; i < len(rooms[room]); i++ {
			if _, ok := m[rooms[room][i]]; !ok {
				m[rooms[room][i]] = struct{}{}
				queue = append(queue, rooms[room][i])
			}
		}
	}
	if len(m) == len(rooms) {
		return true
	}
	return false
}
