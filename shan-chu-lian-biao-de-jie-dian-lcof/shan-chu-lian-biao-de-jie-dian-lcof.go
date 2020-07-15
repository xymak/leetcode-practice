package shan_chu_lian_biao_de_jie_dian_lcof

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {

	now := &ListNode{
		Next: head,
	}
	first := now

	for {
		if now.Next == nil {
			break
		}

		if now.Next.Val == val {
			now.Next = now.Next.Next
			break
		} else {
			now = now.Next
		}
	}

	return first.Next
}
