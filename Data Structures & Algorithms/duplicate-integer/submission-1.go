func hasDuplicate(nums []int) bool {
	seen := make(map[int]bool, len(nums))
	for _, v := range nums {
		if seen[v] {
			return true
		}
		seen[v] = true
	}
	return false
}
