package dnc

func myPow(x float64, n int) float64 {
	if 0 == n {
		return 1.0
	}
	if n < 0 {
		res := myPowPositive(x, -n)
		return 1.0 / res
	}
	return myPowPositive(x, n)
}

func myPowPositive(x float64, n int) float64 {
	if 0 == n {
		return 1.0
	}
	half := myPowPositive(x, n/2)
	res := half * half
	if n%2 == 1 {
		res *= x
	}
	return res
}
