func twoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))
	for i, v := range nums {
		complement := target - v
		if j, ok := seen[complement]; ok {
			return []int{j, i} // j<i since j is found before
		}
		seen[v] = i
	}
	return nil // unreachable
}
