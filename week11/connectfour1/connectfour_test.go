package main

import (
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// the first player is `X`
func TestNewConnectFourBoardFirstPlayerIsX(t *testing.T) {
	board := NewConnectFourBoard()
	assert.Equal(t, 'X', board.player, "initial player should be X")
}

// new boards are empty
func TestNewConnectFourBoardIsEmpty(t *testing.T) {
	board := NewConnectFourBoard()
	// expect board.grid[row][col] == '.' for 6 rows and 7 columns
	for row := range 6 {
		for col := range 7 {
			assert.Equal(t, '.', board.grid[row][col], "every cell should be open")
		}
	}
}

// the board string should have 6 lines of 7 symbols
func TestStringHas6RowsAnd7Columns(t *testing.T) {
	// given a board
	board := NewConnectFourBoard()
	// when i get the string representation of the board
	str := board.String()
	// then the string should have 6 lines
	lines := strings.Split(str, "\n")
	numLines := len(lines)
	require.Equal(t, 6, numLines, "string should have 6 lines")
	// and each line of the string should have 7 runes
	for row := range 6 {
		numRunes := utf8.RuneCountInString(lines[row])
		assert.Equal(t, 7, numRunes, "row %d should have 7 runes", row)
	}
}

// the valid symbols are: '.', 'X', 'O'
func TestStringContainsOnlyValidSymbols(t *testing.T) {
	// given a board with some pieces on it
	board := ConnectFourBoard{
		grid: [][]rune{
			[]rune("......."),
			[]rune(".X....."),
			[]rune("......."),
			[]rune("....O.."),
			[]rune("......."),
			[]rune("......."),
		},
	}
	// when i get the string representation of the board
	str := board.String()
	// then only '.', 'X', and 'O' symbols should be in the string
	open := strings.Count(str, ".")
	x := strings.Count(str, "X")
	o := strings.Count(str, "O")
	assert.Equal(t, 42, open+x+o, "should be 42 '.', 'X', and 'O'")
	assert.Equal(t, 47, len(str), "total length should be 47")
}

// pieces drop to the lowest available empty cell in that column
// multiple pieces in the same column stack upward
// the current player switches after a valid move
func TestPiecesDropToLowestEmptyCell(t *testing.T) {
	// given an empty board
	board := NewConnectFourBoard()
	// when i a drop a piece into the board at column 0
	board.AddPiece(0)
	// then the piece should be at the bottom of the board (last row)
	// and it should be an X
	assert.Equal(t, 'X', board.grid[5][0], "piece should drop to bottom row")

	// when i drop a piece at column 0
	board.AddPiece(0)
	// then the piece should be on to of the existsing piece column 0
	// and it should be an O
	assert.Equal(t, 'O', board.grid[4][0], "O piece should drop to sit on X")
	assert.Equal(t, 'X', board.grid[5][0], "X piece should still be on bottom row")
	// fmt.Println(board)
}

// the current player does not switch after an invalid move
func TestCurrentPlayerDoesNotSwitchOnInvalidMove(t *testing.T) {
	// given a board with a full column
	board := ConnectFourBoard{
		grid: [][]rune{
			[]rune(".O....."),
			[]rune(".X....."),
			[]rune(".O....."),
			[]rune(".X....."),
			[]rune(".O....."),
			[]rune(".X....."),
		},
		player: 'X',
	}
	// when i add a piece to that column
	currentPlayer := board.player
	board.AddPiece(1)
	// then the current player stays the same
	assert.Equal(t, currentPlayer, board.player, "expected player to not change on invalid move")

	// when i add a piece out of bounds
	board.AddPiece(-1)
	// then the current player stays the same
	assert.Equal(t, currentPlayer, board.player, "expected player to not change on invalid move")

	// when i add a piece out of bounds
	board.AddPiece(7)
	// then the current player stays the same
	assert.Equal(t, currentPlayer, board.player, "expected player to not change on invalid move")
}

func TestIsGameOverEmptyBoard(t *testing.T) {
	// given an empty board
	board := NewConnectFourBoard()
	// then the game should not be over
	assert.False(t, board.IsGameOver(), "the game should not be over with an empty board")
}

func TestIsGameOverFullBoard(t *testing.T) {
	// given a full board
	board := ConnectFourBoard{
		grid: [][]rune{
			[]rune("XOXXXOX"),
			[]rune("OXOOOXO"),
			[]rune("OXXXOOO"),
			[]rune("XOOOXXX"),
			[]rune("XOXOXXO"),
			[]rune("XOOOXOX"),
		},
	}
	// then the game should be over
	assert.True(t, board.IsGameOver(), "the game should be over when the board is full")
}

func TestIsGameOverWithWinInARow(t *testing.T) {
	// given a board with 4 in a row
	board := ConnectFourBoard{
		grid: [][]rune{
			[]rune("......."),
			[]rune("......."),
			[]rune("......."),
			[]rune("......."),
			[]rune(".OO...."),
			[]rune("XXXXO.."),
		},
	}
	// then the game should be over
	assert.True(t, board.IsGameOver(), "the game should be over if a player has 4 in a row")
}

func TestIsGameOverWithWinInAColumn(t *testing.T) {
	// given a board with 4 in a column
	board := ConnectFourBoard{
		grid: [][]rune{
			[]rune("......."),
			[]rune("......."),
			[]rune("....X.."),
			[]rune("...OX.."),
			[]rune("...OX.."),
			[]rune("...OX.."),
		},
	}
	// then the game should be over
	assert.True(t, board.IsGameOver(), "the game should be over if a player has 4 in a column")
}

func TestIsGameOverWithWinInADiagonal(t *testing.T) {
	// given a board with 4 in a down-left diagonal
	board := ConnectFourBoard{
		grid: [][]rune{
			[]rune("......."),
			[]rune("......."),
			[]rune("...O..."),
			[]rune("..OX..."),
			[]rune(".OXXO.."),
			[]rune("OXXXO.."),
		},
	}
	// then the game should be over
	assert.True(t, board.IsGameOver(), "the game should be over if a player has 4 in a diagonal")

	// given a board with 4 in a down-right diagonal
	board = ConnectFourBoard{
		grid: [][]rune{
			[]rune("......."),
			[]rune("......."),
			[]rune("...O..."),
			[]rune("...XO.."),
			[]rune("..OXXO."),
			[]rune("..OXXXO"),
		},
	}
	// then the game should be over
	assert.True(t, board.IsGameOver(), "the game should be over if a player has 4 in a diagonal")
}

func TestGetWinner(t *testing.T) {
	// board where X wins
	{
		board := ConnectFourBoard{
			grid: [][]rune{
				[]rune("......."),
				[]rune("......."),
				[]rune("......."),
				[]rune("......."),
				[]rune(".OO...."),
				[]rune("XXXXO.."),
			},
		}
		assert.Equal(t, 'X', board.GetWinner(), "X should be the winner")
	}

	// board where O wins
	{
		board := ConnectFourBoard{
			grid: [][]rune{
				[]rune("......."),
				[]rune("......."),
				[]rune("...O..."),
				[]rune("..OX..."),
				[]rune(".OXXO.."),
				[]rune("OXXXO.."),
			},
		}
		assert.Equal(t, 'O', board.GetWinner(), "O should be the winner")
	}

	// board where column winner
	{
		board := ConnectFourBoard{
			grid: [][]rune{
				[]rune("......."),
				[]rune("......."),
				[]rune("....X.."),
				[]rune("...OX.."),
				[]rune("...OX.."),
				[]rune("...OX.."),
			},
		}
		assert.Equal(t, 'X', board.GetWinner(), "X should be the winner")
	}

	// board where no one wins (incomplete)
	{
		board := NewConnectFourBoard()
		assert.NotEqual(t, 'X', board.GetWinner(), "X is not the winner of an empty board")
		assert.NotEqual(t, 'O', board.GetWinner(), "O is not the winner of an empty board")
	}

	// board where no one wins (draw)
	{
		board := ConnectFourBoard{
			grid: [][]rune{
				[]rune("XXOOOXX"),
				[]rune("XOXXXOO"),
				[]rune("OOXOXXO"),
				[]rune("OXOOOXX"),
				[]rune("OOXXOOO"),
				[]rune("XXOXXOX"),
			},
		}
		assert.NotEqual(t, 'X', board.GetWinner(), "X is not the winner of a drawn board")
		assert.NotEqual(t, 'O', board.GetWinner(), "O is not the winner of a drawn board")
	}
}
