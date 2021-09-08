package main

import (
	"math"
)

func nearestValidPoint(x int, y int, points [][]int) int {
	validIndex := -1
	minDistance := math.MaxInt64
	for index, v := range points {
		if v[0] == x || v[1] == y {
			distance := abs(x, v[0]) + abs(y, v[1])
			if distance < minDistance {
				minDistance = distance
				validIndex = index
			}
		}
	}
	return validIndex
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
