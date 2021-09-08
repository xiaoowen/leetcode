package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := make(map[int]int)

	for _, va := range nums1 {
		for _, vb := range nums2 {
			cnt := va + vb
			if _, ok := m[cnt]; ok {
				m[cnt]++
			} else {
				m[cnt] = 1
			}
		}
	}

	times := 0
	for _, vc := range nums3 {
		for _, vd := range nums4 {
			if i, ok := m[0-(vc+vd)]; ok {
				times += i
			}
		}
	}
	return times
}
