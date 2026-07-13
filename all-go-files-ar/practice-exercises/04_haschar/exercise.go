package main

func HasChar(s string, target rune) bool {
	for _, ch := range s {
		if ch == target {
			return true
		}
	}
	return false
}
