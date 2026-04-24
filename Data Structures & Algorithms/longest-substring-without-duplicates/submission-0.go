func lengthOfLongestSubstring(s string) int {
	// intuition: dynamic sliding window with 'seen' set

	// simple cases
	if len(s) <= 1 { return len(s) }

	seen := make(map[byte]bool)
	longest := 1
	iLeft, iRight := 0, 1
	seen[s[iLeft]] = true
	for iRight < len(s) {
		if seen[s[iRight]] {
			seen[s[iLeft]] = false
			iLeft++
		} else {
			windowSize := iRight - iLeft + 1
			if windowSize > longest { longest = windowSize }
			seen[s[iRight]] = true
			iRight++
		}
	}
	return longest
}
