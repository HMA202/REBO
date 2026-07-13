package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}

	result := 1
	maxInt := int(^uint(0) >> 1)

	for i := 2; i <= nb; i++ {
		if result > maxInt/i {
			return 0
		}

		result *= i
	}

	return result
}
