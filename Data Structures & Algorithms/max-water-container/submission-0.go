func maxArea(heights []int) int {
    // intuition: we need to look for min(height diff) and max(index diff)
    // naive: iterate over the array (combinatoric) -> O(n^2)
    // maybe use two pointer? but how to decide the movement of pointers?
    // maybe move the pointer that has the smaller height?
    // -> true. think of it thath the bottleneck is always on the shorter
    //    side. raising the taller side does nothing, water still leaks
    //    from the otther side. when we move pointer inward, width always
    //    decreased, so to have any chance finding more water, we need
    //    height to increase

    max := 0
    iLeft, iRight := 0, len(heights)-1
    for iLeft < iRight {
        area := (iRight - iLeft) * min(heights[iRight], heights[iLeft])
        if area > max { max = area }

        // now move the shorter side
        if heights[iRight] > heights[iLeft] {
            iLeft++
        } else {
            iRight--
        }
    }
    return max
}
