package main

func findPoisonedDuration(timeSeries []int, duration int) int {
	total := 0

	if len(timeSeries) == 0 {
		return total
	}

	for i := 1; i < len(timeSeries); i++ {
		diff := timeSeries[i] - timeSeries[i-1]
		if diff < duration {
			total += diff
		} else {
			total += duration
		}
	}
	return total + duration
}
