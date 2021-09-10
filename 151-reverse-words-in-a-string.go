package main

import "strings"

func reverseWords(s string) string {
	words := strings.Split(s, " ")

	newWorks := make([]string, 0)
	j := len(words) - 1
	for j >= 0 {
		if words[j] != "" {
			newWorks = append(newWorks, words[j])
		}
		j--
	}

	return strings.Join(newWorks, " ")
}
