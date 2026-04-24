func threeSum(nums []int) [][]int {
    // intuition: fix one number, then search the sum of remaining
    // that equal to the negative of the fixed number
    // -> this requires two pointers that need array to be sorted
    sort.Ints(nums)

    result := [][]int{}
    for iFixed := range len(nums) {
        if iFixed > 0 && nums[iFixed] == nums[iFixed-1] {
            continue // avoid duplicate, using the fact that it is sorted
        }

        target := -1 * nums[iFixed]
        iLeft, iRight := iFixed+1, len(nums)-1
        for iLeft < iRight {
            left, right := nums[iLeft], nums[iRight]
            if left + right < target {
                iLeft++ // if smaller, move iLeft
            } else if left + right > target { 
                iRight-- // if bigger, move iRight
            } else {
                result = append(result, []int{nums[iFixed], left, right})
                // after finding, move both inward to continue searching
                iLeft++
                iRight--
                // avoid duplicates, using the fact that array is sorted
                for iLeft < iRight && nums[iLeft] == nums[iLeft-1] { iLeft++ }
                for iLeft < iRight && nums[iRight] == nums[iRight+1] { iRight-- }
            }
        }
    }
    return result
}
