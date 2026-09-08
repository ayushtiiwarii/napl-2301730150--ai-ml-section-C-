package mathutil

func Power(b, p int) int {
	result := 1

	for i := 0; i < p; i++ {
		result *= b
	}

	return result
}
