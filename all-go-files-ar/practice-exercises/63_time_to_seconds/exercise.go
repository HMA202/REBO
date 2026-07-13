package main

func twoDigitsToInt(a byte, b byte) int {
	return int(a-'0')*10 + int(b-'0')
}

func TimeToSeconds(s string) int {
	if len(s) != 8 {
		return -1
	}

	if s[2] != ':' || s[5] != ':' {
		return -1
	}

	hours := twoDigitsToInt(s[0], s[1])
	minutes := twoDigitsToInt(s[3], s[4])
	seconds := twoDigitsToInt(s[6], s[7])

	if minutes < 0 || minutes > 59 || seconds < 0 || seconds > 59 {
		return -1
	}

	return hours*3600 + minutes*60 + seconds
}
