func productExceptSelf(nums []int) []int {
    prefix := make([]int, len(nums))
    preC := 1
    for i, num := range(nums) {
        preC *= num
        prefix[i] = preC
    }

    suffix := make([]int, len(nums))
    sufC := 1
    for i := len(nums)-1; i >= 0; i-- {
        sufC *= nums[i]
        suffix[i] = sufC
    }

    result := make([]int, len(nums))
    for i := range len(nums) {
        var pre, suf int
        if i-1 < 0 { pre = 1 } else { pre = prefix[i-1] }
        if i+1 > len(nums)-1 { suf = 1 } else { suf = suffix[i+1] }
        result[i] = pre * suf
    }
    return result
}
