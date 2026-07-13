package main

func SumDigitsInString(s string) int {
	sum := 0

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			sum += int(s[i] - '0')
		}
	}

	return sum
}
