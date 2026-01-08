package strategy

import "tic-tac-toe/internal/game"

type DiagonalWinningStrategy struct{}

func (s DiagonalWinningStrategy) CheckWinner(board *game.Board, player *game.Player) bool {
	grid := board.Grid()
	size := board.Size()

	// main diagonal
	all := true
	for i := 0; i < size; i++ {
		if grid[i][i].Symbol != player.Symbol {
			all = false
			break
		}
	}
	if all {
		return true
	}

	// anti diagonal
	all = true
	for i := 0; i < size; i++ {
		if grid[i][size-1-i].Symbol != player.Symbol {
			all = false
			break
		}
	}
	return all
}
