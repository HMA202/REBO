package main

func StringToUint(s string) uint {
	var result uint = 0

	if s == "" {
		return 0
	}

	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}

		result = result*10 + uint(s[i]-'0')
	}

	return result
}
