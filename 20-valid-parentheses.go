package main

func isValid(s string) bool {
	pairs := map[byte]byte{
		'(':')', '{':'}', '[':']',
	}

	stack := make([]byte, 0)
	for _, v := range s {
		if tmp, ok := pairs[byte(v)]; ok {
			stack = append(stack, tmp)
		} else {
			if len(stack) == 0 || stack[len(stack) - 1] != byte(v) {
				return false
			}
			stack = stack[:len(stack) - 1]
		}
	}
	return len(stack) == 0
}
