package main

func RemoveSpaces(s string) string {
	result := ""
	for _, ch := range s {
		if ch != ' ' {
			result += string(ch)
		}
	}
	return result
}
