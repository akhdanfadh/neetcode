func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for {
		curr := numbers[left] + numbers[right]
		if curr > target {
			right--
		} else if curr < target {
			left++
		} else {
			return []int{left+1, right+1}
		}
	}
}
