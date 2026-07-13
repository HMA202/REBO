package main

func DecimalToBinary(n int) string {
	if n == 0 {
		return "0"
	}

	if n < 0 {
		return ""
	}

	result := ""

	for n > 0 {
		result = string(rune('0'+n%2)) + result
		n /= 2
	}

	return result
}
