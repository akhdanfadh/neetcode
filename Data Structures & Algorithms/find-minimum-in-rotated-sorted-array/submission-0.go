func findMin(nums []int) int {
	// intuition: O(logn) means dividing
	// observation: since sorted already, we can check where the min is
	//   by looking at each partition whether the start id and end id is still sorted
	return partition(nums, 0, len(nums)-1)
}

func partition(nums []int, sId, eId int) int {
	if eId-sId+1 <= 1 { return nums[sId] }
	// this part is already sorted, min is most left
	if nums[sId] <= nums[eId] { return nums[sId] }

	// at this point, rotation exists here, check which half
	mId := sId + (eId - sId) / 2
	if nums[sId] > nums[mId] { // min in left half
		return partition(nums, sId, mId)
	} else {
		return partition(nums, mId+1, eId)
	}
}
