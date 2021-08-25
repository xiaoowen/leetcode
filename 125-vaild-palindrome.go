package main

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	// 双指针，一个从左往右，一个从右往左
	i,j := 0, len(s) - 1
	for i < j {
		if !valid(s[i]) {
			i++
			continue
		}
		if !valid(s[j]) {
			j--
			continue
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func valid(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}
