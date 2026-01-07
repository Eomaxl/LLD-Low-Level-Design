package game

import "fmt"

type Board struct {
	size       int
	movesCount int
	grid       [][]Cell
}

func NewBoard(size int) *Board {
	b := &Board{size: size}
	b.initializeBoard()
	return b
}

func (b *Board) initializeBoard() {
	b.grid = make([][]Cell, b.size)
	for r := 0; r < b.size; r++ {
		b.grid[r] = make([]Cell, b.size)
		for c := 0; c < b.size; c++ {
			b.grid[r][c] = Cell{Symbol: Empty}
		}
	}
	b.movesCount = 0
}

func (b *Board) IsFull() bool {
	return b.movesCount == b.size*b.size
}

func (b *Board) PlaceSymbol(r, c int, sym Symbol) error {
	if r < 0 || r > b.size || c < 0 || c > b.size {
		return fmt.Errorf("cell out of bounds : (%d,%d)", r, c)
	}

	if b.grid[r][c].Symbol != Empty {
		return fmt.Errorf("cell already occupied : (%d, %d)", r, c)
	}

	b.grid[r][c].Symbol = sym
	b.movesCount++
	return nil
}

func (b *Board) Grid() [][]Cell {
	return b.grid
}

func (b *Board) Size() int {
	return b.size
}
