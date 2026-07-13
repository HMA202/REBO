package main

func gcdForLCM(a, b uint) uint {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LeastCommonMultiple(a, b uint) uint {
	if a == 0 || b == 0 {
		return 0
	}

	return (a / gcdForLCM(a, b)) * b
}
