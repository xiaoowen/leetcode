package main

func singleNumber(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	for i, v := range m {
		if v == 1 {
			return i
		}
	}
	return 0
}
