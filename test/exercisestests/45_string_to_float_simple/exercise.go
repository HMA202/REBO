package main

func StringToFloatSimple(s string) float64 {
	if s == "" {
		return 0
	}

	sign := 1.0
	index := 0

	if s[0] == '-' {
		sign = -1
		index = 1
	}

	result := 0.0
	decimal := false
	divisor := 10.0

	for index < len(s) {
		ch := s[index]

		if ch == '.' && !decimal {
			decimal = true
			index++
			continue
		}

		if ch < '0' || ch > '9' {
			return 0
		}

		digit := float64(ch - '0')

		if !decimal {
			result = result*10 + digit
		} else {
			result += digit / divisor
			divisor *= 10
		}

		index++
	}

	return result * sign
}
