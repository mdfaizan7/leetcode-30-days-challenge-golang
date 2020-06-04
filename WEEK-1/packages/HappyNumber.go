package packages

// IsHappy is the main function for this problem
func IsHappy(n int) bool {
	for i := 0; i < n; i++ {
		num := 0
		for n > 0 {
			num += square(n % 10)
			n = n / 10
		}

		if num == 1 {
			return true
		}

		n = num
	}

	return false
}

func square(x int) int {
	return x * x
}
