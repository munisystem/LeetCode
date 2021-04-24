package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	sumOfLen := len(nums1) + len(nums2)
	if sumOfLen == 0 {
		return 0.0
	}
	odd := true
	var mi int
	mi = (sumOfLen / 2) + 1
	if sumOfLen%2 == 0 {
		odd = false
	}
	tmp := []int{}
	for len(tmp) < mi {
		if len(nums1) == 0 {
			tmp = append(tmp, nums2[0])
			nums2 = nums2[1:]
			continue
		} else if len(nums2) == 0 {
			tmp = append(tmp, nums1[0])
			nums1 = nums1[1:]
			continue
		}
		n1, n2 := nums1[0], nums2[0]
		if n1 < n2 {
			tmp = append(tmp, n1)
			nums1 = nums1[1:]
		} else {
			tmp = append(tmp, n2)
			nums2 = nums2[1:]
		}
	}
	if odd {
		return float64(tmp[len(tmp)-1])
	} else {
		return float64(tmp[len(tmp)-1]+tmp[len(tmp)-2]) / 2
	}
}
