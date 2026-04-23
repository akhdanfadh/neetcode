// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

func mergeSort(pairs []Pair) []Pair {
	divide(pairs, 0, len(pairs)-1)
	return pairs
}

func divide(pairs []Pair, start, end int) {
	if start >= end { return } // base: 0 or 1 element
	middle := (start + end) / 2 // this will give int
	divide(pairs, start, middle)
	divide(pairs, middle + 1, end)
	conquer(pairs, start, middle, end)
}

func conquer(pairs []Pair, start, middle, end int) {
	left := make([]Pair, middle-start+1)
	right := make([]Pair, end-middle)
	copy(left, pairs[start:middle+1])
	copy(right, pairs[middle+1:end+1])
	iLeft, iRight, iFinal := 0, 0, start
	
	// compare two items across left and right
	for iLeft < len(left) && iRight < len(right) {
		if left[iLeft].Key <= right[iRight].Key {
			pairs[iFinal] = left[iLeft]
			iLeft++
		} else {
			pairs[iFinal] = right[iRight]
			iRight++
		}
		iFinal++
	}

	// either of these will run
	for iLeft < len(left) {
		pairs[iFinal] = left[iLeft]
		iLeft++
		iFinal++
	}
	for iRight < len(right) {
		pairs[iFinal] = right[iRight]
		iRight++
		iFinal++
	}
}