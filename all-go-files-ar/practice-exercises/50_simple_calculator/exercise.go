package main

func SimpleCalculator(expr string) int {
	result := 0
	current := 0
	sign := 1
	hasNumber := false

	for i := 0; i < len(expr); i++ {
		ch := expr[i]

		if ch >= '0' && ch <= '9' {
			current = current*10 + int(ch-'0')
			hasNumber = true
		} else if ch == '+' || ch == '-' {
			if hasNumber {
				result += sign * current
			}

			current = 0
			hasNumber = false

			if ch == '-' {
				sign = -1
			} else {
				sign = 1
			}
		}
	}

	if hasNumber {
		result += sign * current
	}

	return result
}
