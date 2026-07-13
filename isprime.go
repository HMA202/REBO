package piscine

func IsPrime(nb int) bool {
	if nb < 2 {
		return false
	}

	if nb == 2 {
		return true
	}

	if nb%2 == 0 {
		return false
	}

	for divisor := 3; divisor <= nb/divisor; divisor += 2 {
		if nb%divisor == 0 {
			return false
		}
	}

	return true
}
