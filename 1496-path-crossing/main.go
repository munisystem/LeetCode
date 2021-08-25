package main

import "fmt"

const key = "(%d, %d)"

func isPathCrossing(path string) bool {
	var x, y int
	points := map[string]struct{}{}
	points[fmt.Sprintf(key, x, y)] = struct{}{}
	for _, p := range path {
		switch p {
		case 'N':
			y++
		case 'S':
			y--
		case 'E':
			x++
		case 'W':
			x--
		}
		if _, ok := points[fmt.Sprintf(key, x, y)]; ok {
			return true
		}
		points[fmt.Sprintf(key, x, y)] = struct{}{}
	}
	return false
}
