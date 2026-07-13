package piscine

func AtoiBase(s string, base string) int {
	baseRunes := []rune(base)

	if !validAtoiBase(baseRunes) {
		return 0
	}

	result := 0

	for _, r := range s {
		value := runePosition(r, baseRunes)

		if value == -1 {
			return 0
		}

		result = result*len(baseRunes) + value
	}

	return result
}

func validAtoiBase(base []rune) bool {
	if len(base) < 2 {
		return false
	}

	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			return false
		}

		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}

	return true
}

func runePosition(r rune, base []rune) int {
	for i := 0; i < len(base); i++ {
		if r == base[i] {
			return i
		}
	}

	return -1
}
