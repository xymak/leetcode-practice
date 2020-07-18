package zhan_de_ya_ru_dan_chu_xu_lie_lcof

func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int
	var p int
	for _, v := range pushed {
		stack = append(stack, v)

		for {
			if stack[len(stack)-1] == popped[p] {
				p++
				stack = stack[:len(stack)-1]

				if len(stack) == 0 {
					break
				}
			} else {
				break
			}

		}
	}

	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
