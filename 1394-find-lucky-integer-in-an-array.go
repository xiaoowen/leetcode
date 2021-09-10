package main

func findLucky(arr []int) int {
	numCount := make(map[int]int)
	for _, v := range arr {
		numCount[v]++
	}

	luckyNum := -1
	for num, cnt := range numCount {
		if num == cnt && luckyNum < num {
			luckyNum = num
		}
	}
	return luckyNum
}
