func findDuplicate(nums []int) int {
	// 1. fast slow pointer to get the meeting point
	slow, fast := 0, 0
	for {
		fast = nums[nums[fast]]
		slow = nums[slow]
		if slow == fast { break }
	}

	// 2. meeting to cycle entry = head to cycle entry
	slow2 := 0
	for {
		slow = nums[slow]
		slow2 = nums[slow2]
		if slow == slow2 { return slow }
	}
}
