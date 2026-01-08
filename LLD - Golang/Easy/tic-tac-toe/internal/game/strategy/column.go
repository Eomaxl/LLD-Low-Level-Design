package strategy

import "tic-tac-toe/internal/game"

type ColumnWinningStrategy struct{}

func (s ColumnWinningStrategy) CheckWinner(board *game.Board, player *game.Player) bool {
	grid := board.Grid()
	size := board.Size()

	for c := 0; c < size; c++ {
		all := true
		for r := 0; r < size; r++ {
			if grid[r][c].Symbol != player.Symbol {
				all = false
				break
			}
		}
		if all {
			return true
		}
	}
	return false
}
