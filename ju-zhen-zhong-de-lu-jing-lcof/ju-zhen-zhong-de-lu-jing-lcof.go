package ju_zhen_zhong_de_lu_jing_lcof

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[0]); x++ {
			if dfs(board, x, y, word) {
				return true
			}
		}
	}

	return false
}

func dfs(board [][]byte, x int, y int, word string) bool {

	if board[y][x] == word[0] {
		if len(word) == 1 {
			return true
		}

		board[y][x] = []byte("/")[0]
		if x > 0 && board[y][x-1] != board[y][x] && dfs(board, x-1, y, word[1:]) {
			return true
		}

		if y > 0 && board[y-1][x] != board[y][x] && dfs(board, x, y-1, word[1:]) {
			return true
		}

		if x < len(board[0])-1 && board[y][x+1] != board[y][x] && dfs(board, x+1, y, word[1:]) {
			return true
		}

		if y < len(board)-1 && board[y+1][x] != board[y][x] && dfs(board, x, y+1, word[1:]) {
			return true
		}

		board[y][x] = word[0]
	}

	return false
}
