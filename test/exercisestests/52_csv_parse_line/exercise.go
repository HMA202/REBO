package main

func CSVParseLine(line string) []string {
	result := []string{}
	current := ""
	inQuotes := false

	for i := 0; i < len(line); i++ {
		ch := line[i]

		if ch == '"' {
			inQuotes = !inQuotes
		} else if ch == ',' && !inQuotes {
			result = append(result, current)
			current = ""
		} else {
			current += string(ch)
		}
	}

	result = append(result, current)

	return result
}
