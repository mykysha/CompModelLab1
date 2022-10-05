package rightangle

// RightRectangle calculates integral using right angle method.
func RightRectangle(f func(float64) float64, a, b float64, n int) float64 {
	var result float64
	var arg float64
	var h float64

	h = (b - a) / float64(n)

	for i := 1; i <= n; i++ {
		arg = a + float64(i)*h

		result += f(arg)
	}

	return result * h
}
