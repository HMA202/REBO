package main

func CountUpperLower(s string) (int, int) {
	upper := 0
	lower := 0

	for _, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			upper++
		} else if ch >= 'a' && ch <= 'z' {
			lower++
		}
	}

	return upper, lower
}
