package kata

func Century(year int) int {
	return ((year - 1) / 100) + 1
	// or return (year + 99 ) / 100)
}