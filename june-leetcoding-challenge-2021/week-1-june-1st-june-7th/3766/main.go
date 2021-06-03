package main

import (
	"sort"
)

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	mh := horizontalCuts[0]
	if mh < h-horizontalCuts[len(horizontalCuts)-1] {
		mh = h - horizontalCuts[len(horizontalCuts)-1]
	}
	mw := verticalCuts[0]
	if mw < w-verticalCuts[len(verticalCuts)-1] {
		mw = w - verticalCuts[len(verticalCuts)-1]
	}
	for i := 1; i < len(horizontalCuts); i++ {
		if mh < horizontalCuts[i]-horizontalCuts[i-1] {
			mh = horizontalCuts[i] - horizontalCuts[i-1]
		}
	}
	for i := 1; i < len(verticalCuts); i++ {
		if mw < verticalCuts[i]-verticalCuts[i-1] {
			mw = verticalCuts[i] - verticalCuts[i-1]
		}
	}
	return mh * mw % 1000000007
}
