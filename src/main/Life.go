// An implementation of Conway's Game of Life
package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

var source rand.Source
var random rand.Rand

//stores the current state of the grid
type Board struct {
	grid [][]bool
}

//stores height, width, and the Board
//technically dont need to store the height and width, could just use len once we
//make the grid but w/e this is more verbose
type life struct {
	height, width int
	board         Board
}

//checks if the square at x, y will be alive in the next phase
//TODO implement wrapping
func (board *Board) checkAliveNext(x, y, h, w int) bool {
	alive := board.alive(x, y)
	neighbors := 0
	for i := x - 1; i < x+1; i++ {
		for j := y - 1; j < y+1; j++ {
			if i > 0 && i < h && j > 0 && j < h && board.grid[i][j] {
				neighbors += 1
			}
		}
	}
	//using rules from https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life
	if (alive && neighbors < 4 && neighbors > 1) || (!alive && neighbors == 3) {
		return true
	}
	return false
}

//updates every cell with its correct alive state for the next phase
func (board *Board) advance(h, w int) {
	newBoard := createBoard(h, w)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			newBoard.grid[i][j] = board.checkAliveNext(i, j, h, w)
		}
	}
	board.grid = newBoard.grid
}

//creates the initial grid of cells and populates with random alive states
func createBoard(h, w int) Board {
	board := make([][]bool, h)
	for i := 0; i < h; i++ {
		board[i] = make([]bool, w)
		for j := 0; j < w; j++ {
			randomNum := random.Int63n(100)
			//approximate percentage of alive cells can be changed here
			if randomNum < 50 {
				board[i][j] = true
			} else {
				board[i][j] = false
			}
		}
	}
	return Board{board}
}

//returns the state of the cell at h, w
func (board *Board) alive(h, w int) bool {
	return board.grid[h][w]
}

//advances the board forward 1 step
func (l *life) step() {
	l.board.advance(l.height, l.width)
}

// Implements String interface for the Life type
func (l life) String() string {
	var buf bytes.Buffer
	for y := 0; y < l.height; y++ {
		for x := 0; x < l.width; x++ {
			b := byte(' ')
			if l.board.alive(x, y) {
				b = '*'
			}
			buf.WriteByte(b)
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}

//Creates a new random board with the given height and width
func newLife(h, w int) life {
	return life{
		width:  w,
		height: h,
		board:  createBoard(h, w),
	}
}

func Run(size, ticks int) {
	source = rand.NewSource(time.Now().UnixNano())
	random = *rand.New(source)

	l := newLife(size, size)
	fmt.Print("\x0c", l)
	for i := 0; i < ticks; i++ {
		l.step()
		fmt.Print("\x0c", l) // Clear screen and print field.
		time.Sleep(time.Second)
	}
}
