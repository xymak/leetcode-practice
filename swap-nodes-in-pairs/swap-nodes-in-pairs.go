package swap_nodes_in_pairs

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	n := head.Next.Next
	r := head.Next
	r.Next = head
	head.Next = swapPairs(n)

	return r
}
