package main

func CountWords(s string) int {
	count := 0
	inWord := false
	for _, ch := range s {
		if ch == ' ' || ch == '\n' || ch == '\t' {
			inWord = false
		} else if !inWord {
			count++
			inWord = true
		}
	}
	return count
}
