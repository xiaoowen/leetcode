package main

import "strconv"

func dayOfYear(date string) int {
	monthDays := []int{
		0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31,
	}

	// parse
	year, _ := strconv.Atoi(date[0:4])
	month, _ := strconv.Atoi(date[5:7])
	day, _ := strconv.Atoi(date[8:])

	if isLeapYear(year) {
		monthDays[2] = 29
	}

	for i := 1; i < month; i++ {
		day += monthDays[i]
	}
	return day
}

func isLeapYear(year int) bool {
	if year%100 != 0 && year%4 == 0 {
		return true
	}
	if year%100 == 0 && year%400 == 0 {
		return true
	}
	return false
}
