package main

func backspaceCompare(s string, t string) bool {
	sstack := getStack(s)
	tstack := getStack(t)
	if len(sstack) != len(tstack) {
		return false
	}
	for i := 0; i < len(sstack); i++ {
		if sstack[i] != tstack[i] {
			return false
		}
	}
	return true
}

func getStack(s string) []byte {
	stack := make([]byte, 0)
	for _, v := range s {
		char := byte(v)
		if char == '#' {
			// pop
			if len(stack) == 0 {
				continue
			}
			stack = stack[:len(stack) - 1]
		} else {
			// push
			stack = append(stack, char)
		}
	}
	return stack
}
