func pivotIndex(nums []int) int {
    pivot := 0
    leftSum, rightSum := 0, 0
    for _, num := range nums[1:] {
        rightSum += num
    }
    for pivot < len(nums) {
        if leftSum == rightSum {
            return pivot
        }

        leftSum += nums[pivot]
        if pivot == len(nums) - 1 {
            pivot++
        } else {
            pivot++
            rightSum -= nums[pivot]
        }
    }

    return -1
}