package kata

// My Original Kata
/*func MonkeyCount(n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = i+1
	}
	return result
}
*/
//This appends over return declaration
func MonkeyCount(n int) (res []int) {
	for i := 1; i < n+1; i++ {
		res = append(res, i)
	}
	return
}

/*
func monkeyCount(n int) []int {
	result := make([]int, n)

	for i := range result {
		result[i] = i + 1
	}

	return result
}*/