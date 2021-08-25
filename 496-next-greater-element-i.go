package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nextGreaterMap := make(map[int]int)

	stack := make([]int, 0)
	for _, v := range nums2 {
		size := len(stack)
		if size == 0 {
			stack = append(stack, v)
			continue
		}

		for i := 1; i <= size; i++ {
			prev := stack[size-i]
			if v > prev {
				nextGreaterMap[prev] = v
				stack = stack[:size-i]
			} else {
				break
			}
		}
		stack = append(stack, v)
	}

	res := make([]int, len(nums1))
	for i, v := range nums1 {
		res[i] = -1
		if next, ok := nextGreaterMap[v]; ok {
			res[i] = next
		}
	}
	return res
}
