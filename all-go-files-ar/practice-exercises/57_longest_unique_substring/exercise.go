package main

func LongestUniqueSubstring(s string) int {
	lastSeen := [256]int{}

	for i := 0; i < 256; i++ {
		lastSeen[i] = -1
	}

	start := 0
	best := 0

	for i := 0; i < len(s); i++ {
		ch := s[i]

		if lastSeen[ch] >= start {
			start = lastSeen[ch] + 1
		}

		lastSeen[ch] = i

		length := i - start + 1
		if length > best {
			best = length
		}
	}

	return best
}
