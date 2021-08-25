package main

func removeDuplicates(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	slow := 1
	for fast := 1; fast < size; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
