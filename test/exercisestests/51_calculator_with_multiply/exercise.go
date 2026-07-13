package main

func CalculatorWithMultiply(expr string) int {
	total := 0
	last := 0
	num := 0
	op := byte('+')

	for i := 0; i <= len(expr); i++ {
		ch := byte('+')
		if i < len(expr) {
			ch = expr[i]
		}

		if i < len(expr) && ch >= '0' && ch <= '9' {
			num = num*10 + int(ch-'0')
			continue
		}

		if op == '+' {
			total += last
			last = num
		} else if op == '-' {
			total += last
			last = -num
		} else if op == '*' {
			last = last * num
		}

		op = ch
		num = 0
	}

	return total + last
}
