package main

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}

	m := make(map[int]int)

	for i, v := range nums {
		diff := target - v
		if j, ok := m[diff]; ok {
			return []int{j, i}
		}
		m[v] = i
	}
	return nil
}
