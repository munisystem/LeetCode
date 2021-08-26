package main

import "strings"

func isValidSerialization(preorder string) bool {
	nodes := strings.Split(preorder, ",")
	capacity := 1
	for _, node := range nodes {
		capacity--
		if capacity < 0 {
			return false
		}
		if node != "#" {
			capacity += 2
		}
	}
	return capacity == 0
}
