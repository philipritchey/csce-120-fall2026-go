package main

import (
	"slices"
	"strings"
)

type ConnectFourBoard struct {
	grid   [][]rune
	player rune
}

func NewConnectFourBoard() *ConnectFourBoard {
	return &ConnectFourBoard{
		grid: [][]rune{
			slices.Repeat([]rune{'.'}, 7),
			slices.Repeat([]rune{'.'}, 7),
			slices.Repeat([]rune{'.'}, 7),
			slices.Repeat([]rune{'.'}, 7),
			slices.Repeat([]rune{'.'}, 7),
			slices.Repeat([]rune{'.'}, 7),
		},
		player: 'X',
	}
}

func (board *ConnectFourBoard) String() string {
	var lines []string
	for row := range 6 {
		lines = append(lines, string(board.grid[row]))
	}
	return strings.Join(lines, "\n")
}

func (board *ConnectFourBoard) AddPiece(column int) {
	if column < 0 || column > 6 {
		return
	}
	row := 0
	for row < 6 && board.grid[row][column] == '.' {
		row++
	}
	// row >= 6 || grid[row][column] != '.'
	row--
	if row >= 0 {
		board.grid[row][column] = board.player
		if board.player == 'X' {
			board.player = 'O'
		} else {
			board.player = 'X'
		}
	}
}

func (board *ConnectFourBoard) isFull() bool {
	for row := range 6 {
		for col := range 7 {
			if board.grid[row][col] == '.' {
				return false
			}
		}
	}
	return true
}

func (board *ConnectFourBoard) getWinnerInARow() rune {
	for row := range 6 {
		for col := range 4 {
			piece := board.grid[row][col]
			if piece != '.' &&
				board.grid[row][col+1] == piece &&
				board.grid[row][col+2] == piece &&
				board.grid[row][col+3] == piece {
				return piece
			}

		}
	}
	return '.'
}

func (board *ConnectFourBoard) getWinnerInAColumn() rune {
	for col := range 7 {
		for row := range 3 {
			piece := board.grid[row][col]
			if piece != '.' &&
				board.grid[row+1][col] == piece &&
				board.grid[row+2][col] == piece &&
				board.grid[row+3][col] == piece {
				return piece
			}

		}
	}
	return '.'
}

func (board *ConnectFourBoard) getWinnerInADiagonal() rune {
	for row := range 3 {
		for col := range 4 {
			// down-right diagonal
			piece := board.grid[row][col]
			if piece != '.' &&
				board.grid[row+1][col+1] == piece &&
				board.grid[row+2][col+2] == piece &&
				board.grid[row+3][col+3] == piece {
				return piece
			}

			// down-left diagonal
			piece = board.grid[row][col+3]
			if piece != '.' &&
				board.grid[row+1][col+2] == piece &&
				board.grid[row+2][col+1] == piece &&
				board.grid[row+3][col] == piece {
				return piece
			}
		}
	}
	return '.'
}

func (board *ConnectFourBoard) fourInARow() bool {
	return board.getWinnerInARow() != '.'
}

func (board *ConnectFourBoard) fourInAColumn() bool {
	return board.getWinnerInAColumn() != '.'
}

func (board *ConnectFourBoard) fourInADiagonal() bool {
	return board.getWinnerInADiagonal() != '.'
}

func (board *ConnectFourBoard) IsGameOver() bool {
	return board.isFull() ||
		board.fourInARow() ||
		board.fourInAColumn() ||
		board.fourInADiagonal()
}

func (board *ConnectFourBoard) GetWinner() rune {
	rowWinner := board.getWinnerInARow()
	colWinner := board.getWinnerInAColumn()
	diagWinner := board.getWinnerInADiagonal()
	if rowWinner == '.' && colWinner == '.' && diagWinner == '.' {
		return '.'
	}
	if rowWinner != '.' {
		return rowWinner
	}
	if colWinner != '.' {
		return colWinner
	}
	return diagWinner
}

// TODO(student): refactor and document
