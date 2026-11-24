package main

import "fmt"

// Exist reports whether word can be traced in the grid through
// 4-directionally adjacent cells, each cell used at most once.
// Backtracking with in-place path marking.
func Exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	rows := len(board)
	var dfs func(r, c, i int) bool
	dfs = func(r, c, i int) bool {
		if r < 0 || r >= rows || c < 0 || c >= len(board[r]) || board[r][c] != word[i] {
			return false
		}
		if i == len(word)-1 {
			return true
		}
		saved := board[r][c]
		board[r][c] = '#' // off the alphabet: cannot match again on this path
		found := dfs(r+1, c, i+1) || dfs(r-1, c, i+1) || dfs(r, c+1, i+1) || dfs(r, c-1, i+1)
		board[r][c] = saved
		return found
	}
	for r := range board {
		for c := range board[r] {
			if dfs(r, c, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	board := [][]byte{
		[]byte("ABCE"),
		[]byte("SFCS"),
		[]byte("ADEE"),
	}
	fmt.Println(Exist(board, "ABCCED"), Exist(board, "ABCB"))
}
