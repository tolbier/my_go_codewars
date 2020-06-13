package orderedcount

func OrderedCount(text string) []Tuple {
	var slice []*Tuple
	tupleMap := make(map[rune]*Tuple)
	for _, ch := range text {
		if tuple, ok := tupleMap[ch]; ok {
			tuple.Count = tuple.Count + 1
		} else {
			tuple := &Tuple{ch, 1}
			slice = append(slice, tuple)
			tupleMap[ch] = tuple
		}
	}
	result := []Tuple{}
	for _, tuple := range slice {
		result = append(result, *tuple)
	}
	return result
}
