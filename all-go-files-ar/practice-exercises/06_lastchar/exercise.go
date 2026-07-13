package main

func LastChar(s string) rune {
	last := rune(0)
	for _, ch := range s {
		last = ch
	}
	return last
}
