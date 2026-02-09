// Package validsudoku provides utilities to validate Sudoku boards.
package validsudoku

func isValidSudoku(board [][]byte) bool {
	// An Array that has a hashmap[value]=bool with 9 elements.
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	squares := make([]map[byte]bool, 9)
	// Each map at every index needs to be init in Golang.
	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		squares[i] = make(map[byte]bool)
	}
	// You have to use a nested loop.
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			// Skip any empty cells.
			if board[r][c] == '.' {
				continue
			}
			// Single value on the board
			val := board[r][c]
			// This is how you group 3x3 blocks together.
			boxRegionIdx := (r/3)*3 + c/3

			// If this board value is seen at any other time in a row, column, or is a duplicate
			// value in the a 3x3 box region than it's a invalid puzzle.
			if rows[r][val] || cols[c][val] || squares[boxRegionIdx][val] {
				return false
			}
			// Update maps with what was seen.
			rows[r][val] = true
			cols[c][val] = true
			squares[boxRegionIdx][val] = true
		}
	}
	return true
}
