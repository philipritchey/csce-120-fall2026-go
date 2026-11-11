package main

import "strings"

type Board struct {
	grid   [6][7]rune
	player rune
}

func NewBoard() *Board {
	board := Board{}
	for row := range 6 {
		for col := range 7 {
			board.grid[row][col] = '.'
		}
	}
	board.player = 'X'
	return &board
}

func (board *Board) String() string {
	var str strings.Builder
	for row := range 6 {
		for col := range 7 {
			str.WriteRune(board.grid[row][col])
		}
		str.WriteByte('\n')
	}
	return strings.TrimSpace(str.String())
}

func (board *Board) PlacePiece(column int) {
	if column < 0 || column > 6 {
		return
	}
	row := 5
	for row >= 0 && board.grid[row][column] != '.' {
		row--
	}
	// row < 0 || board.grid[row][column] == '.'

	if row < 0 {
		// column is full
		return
	}

	board.grid[row][column] = board.player
	if board.player == 'X' {
		board.player = 'O'
	} else {
		board.player = 'X'
	}
}

func (board *Board) isFull() bool {
	return !strings.Contains(board.String(), ".")
}

func (board *Board) hasFourInARow() bool {
	return false
}

func (board *Board) hasFourInAColumn() bool {
	return false
}

func (board *Board) hasFourInADiagonal() bool {
	return false
}

func (board *Board) hasConnectFour() bool {
	// 4 in a row
	// or 4 in a column
	// or 4 in a diagonal
	return board.hasFourInARow() ||
		board.hasFourInAColumn() ||
		board.hasFourInADiagonal()
}

func (board *Board) IsGameOver() bool {
	return board.isFull() || board.hasConnectFour()
}
