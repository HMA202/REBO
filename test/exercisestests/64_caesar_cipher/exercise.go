package main

func CaesarCipher(s string, shift int) string {
	result := ""
	shift = shift % 26

	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			offset := int(ch - 'a')
			result += string(rune('a' + (offset+shift+26)%26))
		} else if ch >= 'A' && ch <= 'Z' {
			offset := int(ch - 'A')
			result += string(rune('A' + (offset+shift+26)%26))
		} else {
			result += string(ch)
		}
	}

	return result
}
