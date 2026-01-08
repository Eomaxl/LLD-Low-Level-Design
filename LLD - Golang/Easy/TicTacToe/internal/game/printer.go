package game

import "fmt"

func (b *Board) PrintBoard() {
	for r := 0; r < b.size; r++ {
		for c := 0; c < b.size; c++ {
			fmt.Printf("%c", b.grid[r][c].Symbol.Char())
		}
		fmt.Println()
	}
}
