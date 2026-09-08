package mathutil

func Factorial(k int) int {
	if k < 0 {
		return 0
	}
	result := 1

	for i := 1; i <= k; i++ {
		result *= i
	}
	return result
}
