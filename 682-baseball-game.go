package main

import "strconv"

func calPoints(ops []string) int {
	pointsRecord := make([]int, 0)

	for _, op := range ops {
		size := len(pointsRecord)
		switch {
		case op == "C" && size >= 1:
			pointsRecord = pointsRecord[:size-1]
		case op == "D" && size >= 1:
			pointsRecord = append(pointsRecord, 2*pointsRecord[size-1])
		case op == "+" && size >= 2:
			pointsRecord = append(pointsRecord, pointsRecord[size-1]+pointsRecord[size-2])
		default:
			points, _ := strconv.Atoi(op)
			pointsRecord = append(pointsRecord, points)
		}
	}

	total := 0
	for _, points := range pointsRecord {
		total += points
	}
	return total
}
