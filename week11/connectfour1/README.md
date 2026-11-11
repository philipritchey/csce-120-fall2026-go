# Connect Four Engine
## Requirements
### Can create a new Connect Four board
1. new boards are empty
2. the first player is `X`

### Can drop pieces into/onto a board in columns
1. pieces drop to the lowest available empty cell in that column
2. multiple pieces in the same column stack upward
3. the current player switches after a valid move
4. the current player does not switch after an invalid move:
   - column is full
   - column is out of bounds

### Can determine the state of a board
1. a board with 4 `X` in a "row" (or column or diagonal) is game over and `X` wins
1. a board with 4 `O` in a "row" is game over and `O` wins
1. a full board with no 4 in "row" is game over and a tie
1. a board that is not full and has no 4 in a "row" is incomplete/unfinished

### Can represent a board as a string
1. the board string should have 6 lines of 7 symbols
   1. the valid symbols are:
      - open spaces: `.`
      - player 1's pieces: `X`
      - player 2's pieces: `O`
2. the pieces should be in the correct places

## Design
### Data representation
* Connect Four Board (`type ConnectFourBoard struct`)
  - `grid`: grid of 6 rows and 7 columns (probably a `[][]rune`)
  - `player`: current player, either `'X'` or `'O'` (`rune`)

### Functions and methods
* `func NewConnectFourBoard() *ConnectFourBoard`: create a new board
* `func (board *ConnectFourBoard) AddPiece(column)`: add piece to board and update current player on valid move
* `func (board *ConnectFourBoard) IsGameOver() bool`: determine if game is over
* `func (board *ConnectFourBoard) GetWinner() rune`: determine winner
* `func (board *ConnectFourBoard) String() string`: board as string
