package piscine

func NRune(s string, n int) rune {
	if n <= 0 {
		return 0
	}

	position := 1

	for _, r := range s {
		if position == n {
			return r
		}

		position++
	}

	return 0
}
