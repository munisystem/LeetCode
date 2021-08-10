package main

import (
	"sort"
)

func findRestaurant(list1 []string, list2 []string) []string {
	m := map[string]int{}
	for i := 0; i < len(list1); i++ {
		m[list1[i]] = i
	}

	type kv struct {
		Key   string
		Value int
	}
	var kvs []kv
	for i := 0; i < len(list2); i++ {
		j, ok := m[list2[i]]
		if !ok {
			continue
		}
		kvs = append(kvs, kv{Key: list2[i], Value: i + j})
	}
	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].Value < kvs[j].Value
	})

	var ans []string
	var min int
	for i := 0; i < len(kvs); i++ {
		if i == 0 {
			min = kvs[i].Value
		}
		if min < kvs[i].Value {
			break
		}
		ans = append(ans, kvs[i].Key)
	}
	return ans
}
