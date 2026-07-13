package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}

	if nb%2 == 0 {
		nb++
	}

	for !isPrimeNext(nb) {
		nb += 2
	}

	return nb
}

func isPrimeNext(nb int) bool {
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
