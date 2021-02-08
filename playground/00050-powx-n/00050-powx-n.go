package _00050_powx_n

func myPow(x float64, n int) float64 {
	if n == 0 { // special case
		return 1
	}

	if n == 1 { // special case
		return x
	}

	if n < 0 { // if n ltz just change x
		n = -n
		x = 1 / x
	}
	tmp := myPow(x, n/2)
	if n%2 == 0 { // even number
		return tmp * tmp
	}
	return tmp * tmp * x // odd number
}
