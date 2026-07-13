package main

func StringToInt(s string) int {
	if s == "" {
		return 0
	}

	sign := 1
	index := 0
	result := 0

	if s[0] == '-' {
		sign = -1
		index = 1
	} else if s[0] == '+' {
		index = 1
	}

	for index < len(s) {
		if s[index] < '0' || s[index] > '9' {
			return 0
		}

		result = result*10 + int(s[index]-'0')
		index++
	}

	return result * sign
}
