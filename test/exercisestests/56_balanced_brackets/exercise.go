package main

func BalancedBrackets(s string) bool {
	stack := []rune{}

	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
		} else if ch == ')' || ch == ']' || ch == '}' {
			if len(stack) == 0 {
				return false
			}

			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if ch == ')' && last != '(' {
				return false
			}
			if ch == ']' && last != '[' {
				return false
			}
			if ch == '}' && last != '{' {
				return false
			}
		}
	}

	return len(stack) == 0
}
