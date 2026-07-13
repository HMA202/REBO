package piscine

func Capitalize(s string) string {
	result := ""
	startOfWord := true

	for _, r := range s {
		isLower := r >= 'a' && r <= 'z'
		isUpper := r >= 'A' && r <= 'Z'
		isDigit := r >= '0' && r <= '9'
		isAlphaNumeric := isLower || isUpper || isDigit

		if isAlphaNumeric {
			if startOfWord {
				if isLower {
					r = r - ('a' - 'A')
				}

				startOfWord = false
			} else if isUpper {
				r = r + ('a' - 'A')
			}
		} else {
			startOfWord = true
		}

		result += string(r)
	}

	return result
}
