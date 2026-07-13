package piscine

func IsPrintable(s string) bool {
	for _, r := range s {
		if r < ' ' || r > '~' {
			return false
		}
	}

	return true
}
