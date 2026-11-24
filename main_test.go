package main

import "testing"

func grid(rows ...string) [][]byte {
	g := make([][]byte, len(rows))
	for i, r := range rows {
		g[i] = []byte(r)
	}
	return g
}

func TestExist(t *testing.T) {
	board := grid("ABCE", "SFCS", "ADEE")
	cases := []struct {
		word string
		want bool
	}{
		{"ABCCED", true},
		{"SEE", true},
		{"ABCB", false}, // would reuse the B
		{"A", true},
		{"XYZ", false},
	}
	for _, c := range cases {
		if got := Exist(board, c.word); got != c.want {
			t.Errorf("Exist(%q) = %v, want %v", c.word, got, c.want)
		}
	}
	// The snaking example needs the board whose middle row is SFES.
	snake := grid("ABCE", "SFES", "ADEE")
	if !Exist(snake, "ABCESEEEFS") {
		t.Error("ABCESEEEFS should snake through the SFES board")
	}
	// The board must be restored after every search.
	if string(board[0]) != "ABCE" || string(board[1]) != "SFCS" || string(board[2]) != "ADEE" {
		t.Error("board was not restored")
	}
	if !Exist(grid("a"), "a") {
		t.Error("single cell match")
	}
	if Exist(grid(), "a") {
		t.Error("empty board cannot contain a word")
	}
}
