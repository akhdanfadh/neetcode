func permute(nums []int) [][]int {
	perms := [][]int{{}}
	for _, num := range nums {
		newPerms := [][]int{}
		for _, perm := range perms {
			for i := range len(perm) + 1 {
                next := make([]int, 0, len(perm)+1)
				next = append(next, perm[:i]...) // bagian sebelum posisi i
				next = append(next, num)         // sisipkan num
				next = append(next, perm[i:]...) // bagian setelah posisi i
				newPerms = append(newPerms, next)// pas ditambahin
			}
		}
		perms = newPerms
	}
	return perms
}
