package main

func WordFrequency(s string) map[string]int {
	frequency := map[string]int{}
	word := ""

	for i := 0; i <= len(s); i++ {
		ch := byte(' ')

		if i < len(s) {
			ch = s[i]
		}

		if ch == ' ' {
			if word != "" {
				frequency[word]++
				word = ""
			}
		} else {
			word += string(ch)
		}
	}

	return frequency
}
