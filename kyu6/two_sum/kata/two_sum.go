package kata

func TwoSum(numbers []int, target int) [2]int {
	m := make(map[int]int)
	for idx, num := range numbers {
		other := target - num
		if val, ok := m[other]; ok {
			return [...]int{val, idx}
		}
		m[num] = idx
	}
	return [2]int{}
}
