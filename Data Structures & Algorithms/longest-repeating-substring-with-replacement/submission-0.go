import "slices"

func characterReplacement(s string, k int) int {
	// intuition: count chars in the window, max char count is the fixed,
	//   so k >= total count (window length) - max char count to replace
	//   the non-fixed chars
	// challenge: how to track the max char count?
	//   maybe instead of map of char to count, we use fixed array like bucket?

	result := 0
	charCount := [26]int{}
	lId := 0
	for rId := 0; rId < len(s); rId++ {
		charCount[s[rId]-'A']++
		if k < (rId-lId+1) - slices.Max(charCount[:]) {
			charCount[s[lId]-'A']--
			lId++
		}
		if result < rId-lId+1 {
			result = rId-lId+1
		}
	}
	return result
}