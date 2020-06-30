package largest_rectangle_in_histogram

type StackElement struct {
	Height int
	Index  int
}

func largestRectangleArea(heights []int) int {
	l := len(heights)
	stack := make([]StackElement, 0, l)
	left := make([]int, l, l)
	right := make([]int, l, l)

	for i := 0; i < l; i++ {
		for {
			if len(stack) == 0 {
				stack = append(stack, StackElement{
					Height: heights[i],
					Index:  i,
				})
				left[i] = -1
				break
			} else if stack[len(stack)-1].Height < heights[i] {
				left[i] = stack[len(stack)-1].Index
				stack = append(stack, StackElement{
					Height: heights[i],
					Index:  i,
				})
				break
			} else {
				stack = stack[:len(stack)-1]
			}

		}
	}

	stack = make([]StackElement, 0, l)

	for i := l - 1; i >= 0; i-- {
		for {
			if len(stack) == 0 {
				stack = append(stack, StackElement{
					Height: heights[i],
					Index:  i,
				})
				right[i] = l
				break
			} else if stack[len(stack)-1].Height < heights[i] {
				right[i] = stack[len(stack)-1].Index
				stack = append(stack, StackElement{
					Height: heights[i],
					Index:  i,
				})
				break
			} else {
				stack = stack[:len(stack)-1]
			}

		}
	}

	max := 0
	for i := 0; i < l; i++ {
		h := heights[i]
		w := right[i] - left[i] - 1

		square := h * w
		if square > max {
			max = square
		}
	}

	return max
}
