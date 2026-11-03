func firstUniqChar(s string) int {
	// count chars
	counter := make(map[rune]int, len(s))
	for _, char := range s {
		counter[char]++
	}
	
	// iterate again to find the first unique
	for idx, char := range s {
		if counter[char] == 1 {
			return idx
		}
	}
	return -1
}
