package main

func anagramKey(word string) string {
	counts := [26]int{}

	for _, ch := range word {
		if ch >= 'a' && ch <= 'z' {
			counts[ch-'a']++
		}
	}

	key := ""

	for i := 0; i < 26; i++ {
		key += string(rune('a' + i))
		key += string(rune('0' + counts[i]))
	}

	return key
}

func GroupAnagrams(words []string) map[string][]string {
	groups := map[string][]string{}

	for _, word := range words {
		key := anagramKey(word)
		groups[key] = append(groups[key], word)
	}

	return groups
}
