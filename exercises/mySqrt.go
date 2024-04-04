package exercises

func MySqrt(x int) int {

	res := 2
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	if x == 3 {
		return 1
	}
	for {
		if (res * res) < x {
			res++
		} else if (res * res) == x {
			return res
		} else {
			return res - 1
		}
	}

}
