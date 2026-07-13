package piscine

func ToLower(s string) string {
	result := ""

	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			r = r + ('a' - 'A')
		}

		result += string(r)
	}

	return result
}
