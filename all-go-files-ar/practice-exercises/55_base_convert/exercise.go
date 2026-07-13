package main

func digitValue(ch byte) int {
	if ch >= '0' && ch <= '9' {
		return int(ch - '0')
	}
	if ch >= 'A' && ch <= 'F' {
		return int(ch-'A') + 10
	}
	if ch >= 'a' && ch <= 'f' {
		return int(ch-'a') + 10
	}

	return -1
}

func digitChar(n int) byte {
	if n < 10 {
		return byte('0' + n)
	}
	return byte('A' + n - 10)
}

func BaseConvert(s string, fromBase int, toBase int) string {
	if fromBase < 2 || fromBase > 16 || toBase < 2 || toBase > 16 {
		return ""
	}

	decimal := 0

	for i := 0; i < len(s); i++ {
		value := digitValue(s[i])
		if value < 0 || value >= fromBase {
			return ""
		}

		decimal = decimal*fromBase + value
	}

	if decimal == 0 {
		return "0"
	}

	result := ""

	for decimal > 0 {
		result = string(digitChar(decimal%toBase)) + result
		decimal /= toBase
	}

	return result
}
