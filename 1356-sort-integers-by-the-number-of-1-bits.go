package main

import (
	"math/bits"
	"sort"
)

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		x, y := arr[i], arr[j]
		cnti, cntj := bits.OnesCount(uint(x)), bits.OnesCount(uint(y))
		return cnti < cntj || (cnti == cntj && x < y)
	})
	return arr
}
