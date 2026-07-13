package main

func ToUpperSimple(s string) string {
	result := ""
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			ch = ch - 32
		}
		result += string(ch)
	}
	return result
}
