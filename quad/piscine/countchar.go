package piscine

func CountChar(s string, target rune) int {
	count := 0

	for _, ch := range s {
		if ch == target {
			count++
		}
	}

	return count
}
