# Connect Four Game Engine
## Requirements
### Can create a new Connect Four board
1. new boards are empty (only open '.' cells)
2. the first player is `X`

### Can drop pieces into/onto a board in columns
1. pieces drop to the lowest available empty cell in that column
2. multiple pieces in the same column stack upward
3. the current player switches after a valid move
4. the current player does not switch after an invalid move:
   - column is full
   - column is out of bounds
5. only `X` or `O` pieces are allowed

### Can determine the state of a board
1. a board with 4 `X` in a "row" (or column or diagonal) is game over and `X` wins
1. a board with 4 `O` in a "row" (or column or diagonal) is game over and `O` wins
1. a board with no empty spaces and no 4 in a "row" is game over and is a drawn game
1. a board with some empty space and no 4 in a "row" is incomplete/unfinished

### Can represent a board as a string
1. the string should have 6 lines of 7 symbols each
   * the valid symbols are:
     - open: '.'
     - player 1's pieces: 'X'
     - player 2's pieces: 'O'
   * no spaces between symbols in the same line

## Design
### Data representation
* `Board`: the connect four board
  - `grid`: grid (probably `[6][7]rune`)
  - `player`: the current player (probably `rune`)

### Functions/methods
* add a piece to the board: `func (*Board) PlacePiece(column)`
  - returns updated board
* makes a new board: `func NewBoard() *Board`
* determine is board is in a winning state: `func (Board) GetState() (bool, rune)`
  - bool: is game over? `func (Board) IsGameOver() bool`
  - rune: winner, if game is over; don't care if game not over `func (Board) GetWinner() rune`
* constructs a string representation of the board: `func (Board) String() string`
