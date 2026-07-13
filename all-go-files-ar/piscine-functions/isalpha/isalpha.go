package piscine

func IsAlpha(s string) bool {
	for _, r := range s {
		isLower := r >= 'a' && r <= 'z'
		isUpper := r >= 'A' && r <= 'Z'
		isDigit := r >= '0' && r <= '9'

		if !isLower && !isUpper && !isDigit {
			return false
		}
	}

	return true
}
