package er_cha_sou_suo_shu_de_hou_xu_bian_li_xu_lie_lcof

func verifyPostorder(postorder []int) bool {
	l := len(postorder)
	if l <= 1 {
		return true
	}

	root := postorder[l-1]

	m := 0
	n := 0
	for i := l - 2; i >= 0; i-- {
		if postorder[i] > root {
			m++
			if n > 0 {
				return false
			}
		} else {
			n++
		}
	}

	left := true
	right := true
	if m > 0 {
		right = verifyPostorder(postorder[l-1-m : l-1])
	}
	if n > 0 {
		left = verifyPostorder(postorder[l-1-m-n : l-1-m])
	}

	return left && right
}
