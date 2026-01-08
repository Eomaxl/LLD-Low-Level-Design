package strategy

import "tic-tac-toe/internal/game"

type RowWinningStrategy struct{}

func (c RowWinningStrategy) CheckWinner(board *game.Board, player *game.Player) bool {
	grid := board.Grid()
	size := board.Size()

	for r := 0; r < size; r++ {
		all := true
		for c := 0; c < size; c++ {
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
