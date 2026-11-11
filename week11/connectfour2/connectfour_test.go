package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// new boards are empty
func TestNewBoardIsEmpty(t *testing.T) {
	// given a new board
	board := NewBoard()
	require.NotNil(t, board, "new board should not be nil")
	// then it should be empty: all elements of grid should be '.' (open)
	for row := range 6 {
		for col := range 7 {
			assert.Equal(t, '.', board.grid[row][col], "initial board should be empty")
		}
	}

}

// the first player is `X`
func TestNewBoardCurrentPlayerIsX(t *testing.T) {
	// given a new board
	board := NewBoard()
	// then the current player should be 'X'
	assert.Equal(t, 'X', board.player, "initial current player should be X")
}

// the string representation of the board should have 6 lines of 7 symbols each
func TestStringHas6LinesOf7Symbols(t *testing.T) {
	// given a board
	board := NewBoard()
	// when i get the string representation of the board
	str := board.String()
	// then the string should have 6 lines
	lines := strings.Split(str, "\n")
	assert.Equal(t, 6, len(lines), "string should have 6 lines")
	// and each line should have 7 characters/runes
	for i, line := range lines {
		numChars := len(strings.Split(line, ""))
		assert.Equal(t, 7, numChars, "string should have 7 characters in line %d", i)
	}
}

// the string representation should only contains symbols: '.', 'X', 'O' (and '\n')
func TestStringContainsOnlyValidSymbols(t *testing.T) {
	// given a board
	board := NewBoard()
	// when i get the string representation of the board
	str := board.String()
	// then the string should contain only '.', 'X',' 'O', and '\n'
	for _, r := range str {
		assert.Contains(t, []rune(".XO\n"), r, "only ., X, O, and newline are allowed")
	}
}

// pieces drop to the lowest available empty cell in that column
// multiple pieces in the same column stack upward
func TestPlacedPieceDropsToLowestAvailableCell(t *testing.T) {
	// given a new board with an X in column 0
	board := NewBoard()
	board.PlacePiece(0)
	// when i place a piece into column 0
	board.PlacePiece(0)
	// then the bottom row of column 0 should have 'X'
	// and the row above it should have 'O'
	assert.Equal(t, 'X', board.grid[5][0], "expected placed piece to fall to bottom row")
	assert.Equal(t, 'O', board.grid[4][0], "expected placed piece to fall to lowest open cell")
}

// the current player switches after a valid move
func TestPlacePieceSwitchesPlayerAfterValidMove(t *testing.T) {
	// given a new board
	board := NewBoard()
	// when i place a piece
	board.PlacePiece(0)
	// then the current player should update
	assert.Equal(t, 'O', board.player, "player should update after valid move")

	// when i place another piece
	board.PlacePiece(0)
	// then the current player should update
	assert.Equal(t, 'X', board.player, "player should update after valid move")
}

// the current player does not switch after an invalid move:
//   - column is full
//   - column is out of bounds
func TestPlacePieceDoesNotUpdatePlayerAfterInvalidMove(t *testing.T) {
	// check out of bounds
	{
		// given a new board
		board := NewBoard()
		// when i place a piece in column -1
		currentPlayer := board.player
		board.PlacePiece(-1)
		// current player should be the same
		assert.Equal(t, currentPlayer, board.player, "player should not update after invalid move")

		// when i place a piece in column 7
		board.PlacePiece(7)
		// current player should be the same
		assert.Equal(t, currentPlayer, board.player, "player should not update after invalid move")
	}

	// check column is full
	{
		// given a board with a full column 1
		var board Board
		for row := range 6 {
			board.grid[row][1] = 'X'
		}
		board.player = 'X'

		// when i place a piece in column 1
		currentPlayer := board.player
		board.PlacePiece(1)
		// then the current player should be the same
		assert.Equal(t, currentPlayer, board.player, "player should not update after invalid move (full column)")
	}
}

// a board with some empty space and no 4 in a "row" is incomplete/unfinished
func TestIsGameOverForNewBoard(t *testing.T) {
	// given a new board
	board := NewBoard()
	// then the game is not over
	assert.False(t, board.IsGameOver(), "the game is not over when the board is empty")
}

// a board with no empty spaces is game over
func TestIsGameOverForFullBoard(t *testing.T) {
	// given a full board
	var board Board
	for row := range 6 {
		for col := range 7 {
			board.grid[row][col] = 'X'
		}
	}
	// then the game is over
	assert.True(t, board.IsGameOver(), "the game is over when the board is full")
}

// a board with 4 in a "row" (or column or diagonal) is game over
func TestIsGameOverWhen4InARow(t *testing.T) {
	// given a board with 4 in a row
	board := Board{
		grid: [6][7]rune{
			{'.', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.'},
			{'X', 'X', 'X', 'X', '.', '.', '.'},
		},
		player: 'X',
	}
	// then the game is over
	assert.True(t, board.IsGameOver(), "the game is over when a player has connect four")
}
