package main

func HexDigitValue(ch rune) int {
	if ch >= '0' && ch <= '9' {
		return int(ch - '0')
	}

	if ch >= 'a' && ch <= 'f' {
		return int(ch-'a') + 10
	}

	if ch >= 'A' && ch <= 'F' {
		return int(ch-'A') + 10
	}

	return -1
}
