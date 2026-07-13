package main

func FirstChar(s string) rune {
	for _, ch := range s {
		return ch
	}
	return 0
}
