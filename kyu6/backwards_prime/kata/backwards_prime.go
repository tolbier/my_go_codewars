package kata

func BackwardsPrime(start int, stop int) []int {
	var result []int

	for i := start; i <= stop; i++ {
		r := Reverse(i)
		if r != i && IsPrime(i) && IsPrime(r) {
			result = append(result, i)
		}
	}
	return result
}

func Reverse(n int) int {
	var r int
	for ; n > 0; n /= 10 {
		r = r*10 + (n % 10)
	}
	return r
}

func IsPrime(n int) bool {
	for m := 2; m*m <= n; m++ {
		if n%m == 0 {
			return false
		}
	}
	return true
}
