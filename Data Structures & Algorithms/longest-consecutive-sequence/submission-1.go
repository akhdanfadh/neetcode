func longestConsecutive(nums []int) int {
    // intuition: need something that indicate start of sequence
    // -> num is sequence starter if num-1 is not in nums
    // so we need to make nums to be set for easier lookup O(1)

    // edge case
    if len(nums) == 0 { return 0 }

    set := map[int]struct{}{}
    for _, num := range nums {
        set[num] = struct{}{}
    }

    starter := []int{}
    for num := range set {
        if _, exists := set[num-1]; !exists {
            starter = append(starter, num)
        }
    }

    // now check over starter
    max := 1
    for _, start := range starter {
        curInt := start
        count := 1
        for {
            if _, exists := set[curInt+1]; !exists {
                break
            }
            count++
            curInt++
        }
        if count > max { max = count }
    }
    return max
}
