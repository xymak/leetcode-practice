package cong_wei_dao_tou_da_yin_lian_biao_lcof

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	} else {
		r := &[]int{}
		dfs(head, r)
		return *r
	}
}

func dfs(head *ListNode, r *[]int) {
	if head.Next != nil {
		dfs(head.Next, r)
	}
	*r = append(*r, head.Val)
}
