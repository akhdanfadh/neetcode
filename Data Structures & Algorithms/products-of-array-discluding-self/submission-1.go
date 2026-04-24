func productExceptSelf(nums []int) []int {
    // store the prefix directly in result
    result := make([]int, len(nums))
    result[0] = 1
    for i := 0; i <= len(nums)-2; i++ {
        // shifted because the result's position
        // is previous prefix * after suffix
        result[i+1] = result[i] * nums[i]
    }

    // now do the suffix, shifted as well
    sufC := 1
    for i := len(nums)-1; i >= 0; i-- {
        result[i] *= sufC
        sufC *= nums[i]
    }
    return result
}
