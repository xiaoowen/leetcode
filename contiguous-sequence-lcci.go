package main

func maxSubArray(nums []int) int {
	numssize := len(nums)
	if numssize == 1 {
		return nums[0]
	}
	sum := nums[0]
	maxsum := nums[0]

	for i := 1; i < numssize; i++ {
		if sum+nums[i] <= nums[i] {
			sum = nums[i]
		} else {
			sum += nums[i]
		}

		if sum > maxsum {
			maxsum = sum
		}
	}
	return maxsum
}
