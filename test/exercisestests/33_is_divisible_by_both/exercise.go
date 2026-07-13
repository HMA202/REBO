package main

func IsDivisibleByBoth(n, a, b uint) bool {
	if a == 0 || b == 0 {
		return false
	}

	return n%a == 0 && n%b == 0
}
