package main

func NextDivisibleByBoth(start, a, b uint) uint {
	if a == 0 || b == 0 {
		return 0
	}

	for n := start; ; n++ {
		if n%a == 0 && n%b == 0 {
			return n
		}
	}
}
