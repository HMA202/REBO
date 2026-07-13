package main

func BinaryToDecimal(s string) int {
	result := 0

	if s == "" {
		return 0
	}

	for i := 0; i < len(s); i++ {
		if s[i] != '0' && s[i] != '1' {
			return 0
		}

		result = result*2 + int(s[i]-'0')
	}

	return result
}
