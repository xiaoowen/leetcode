package main

func isPalindromeNumber(x int) bool {
	if x < 0 {
		return false
	}

	cur := x
	ans := 0

	for x > 0 {
		ans = ans*10 + x%10
		x /= 10
	}

	return cur == ans
}
