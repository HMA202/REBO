package main

func romanValue(ch byte) int {
	if ch == 'I' {
		return 1
	}
	if ch == 'V' {
		return 5
	}
	if ch == 'X' {
		return 10
	}
	if ch == 'L' {
		return 50
	}
	if ch == 'C' {
		return 100
	}
	if ch == 'D' {
		return 500
	}
	if ch == 'M' {
		return 1000
	}

	return 0
}

func RomanToInt(s string) int {
	total := 0

	for i := 0; i < len(s); i++ {
		current := romanValue(s[i])
		next := 0

		if i+1 < len(s) {
			next = romanValue(s[i+1])
		}

		if current < next {
			total -= current
		} else {
			total += current
		}
	}

	return total
}
