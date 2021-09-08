package main

func toLowerCase(s string) string {
	s1 := []rune(s)
	for i, v := range s {
		if v >= 'A' && v <= 'Z' {
			s1[i] += 'a' - 'A'
		}
	}
	return string(s1)
}
