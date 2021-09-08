package main

func largestAltitude(gain []int) int {
	max, high := 0, 0
	for _, v := range gain {
		high += v
		if high > max {
			max = high
		}
	}
	return max
}
