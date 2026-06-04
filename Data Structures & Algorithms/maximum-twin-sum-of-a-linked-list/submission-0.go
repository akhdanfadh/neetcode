/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
	// 1. slow and fast pointer to get the middle
	slow, fast := head, head
	for fast != nil { // even length
		fast = fast.Next.Next
		slow = slow.Next
	}
	// now slow is the start of the 2nd half

	// 2. reverse 2nd half (reverse linked list problem)
	prev, cur := (*ListNode)(nil), slow
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev, cur = cur, next
	}
	// now prev is the tail of the linked list

	// 3. traverse 1st half and the reverse 2nd half
	result := 0
	first, second := head, prev
	for second != nil {
		result = max(result, first.Val+second.Val)
		first, second = first.Next, second.Next
	}
	return result
}