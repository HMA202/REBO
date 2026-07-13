package main

func IntToStringRLE(n int) string {
	if n == 0 {
		return "0"
	}

	result := ""

	for n > 0 {
		result = string(rune('0'+n%10)) + result
		n /= 10
	}

	return result
}

func RunLengthEncode(s string) string {
	result := ""
	i := 0

	for i < len(s) {
		if s[i] == ' ' {
			result += " "
			i++
			continue
		}

		count := 1

		for i+count < len(s) && s[i+count] == s[i] {
			count++
		}

		result += IntToStringRLE(count) + string(s[i])
		i += count
	}

	return result
}
