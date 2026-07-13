package piscine

func ToUpper(s string) string {
	result := ""

	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			r = r - ('a' - 'A')
		}

		result += string(r)
	}

	return result
}
