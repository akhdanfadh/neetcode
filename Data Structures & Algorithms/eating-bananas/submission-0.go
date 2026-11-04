import "slices"

func minEatingSpeed(piles []int, h int) int {
	// binary search over 1 to max in piles as eating rate k
	left, right := 1, slices.Max(piles)
	for left < right {
		mid := left + (right - left) / 2
		if hoursNeeded(piles, mid) <= h { // condition
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left // minimum k
}

func hoursNeeded(piles []int, k int) int {
	total := 0
	for _, pile := range piles {
		total += (pile + k - 1) / k // avoid using float
	}
	return total
}
