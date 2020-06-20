package kata

import (
	"math"
)

func Solve(n int) int {
	for b := int(Sqrt(n)) + 1; Pow2(b)-Pow2(b-1) <= n; b++ {
		a := Sqrt(Pow2(b) - n)
		if intA, ok := Integer(a); ok {
			return Pow2(intA)
		}
	}
	return -1
}
func Integer(n float64) (int, bool) {
	intN := int(n)
	return intN, n == float64(intN)
}
func Sqrt(n int) float64 {
	return math.Sqrt(float64(n))
}
func Pow2(n int) int {
	return n * n
}
