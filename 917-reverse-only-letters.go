package main

func reverseOnlyLetters(s string) string {
	s1 := []byte(s)
	i, j := 0, len(s1)-1
	for i < j {
		if !isLetter(s1[i]) {
			i++
			continue
		}
		if !isLetter(s1[j]) {
			j--
			continue
		}
		s1[i], s1[j] = s1[j], s1[i]
		i++
		j--
	}
	return string(s1)
}

func isLetter(v byte) bool {
	return v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z'
}
