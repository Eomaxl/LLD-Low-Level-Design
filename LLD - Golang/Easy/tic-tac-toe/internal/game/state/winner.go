package state

import (
	"errors"
	"tic-tac-toe/internal/game"
)

type WinnerState struct{}

func (st WinnerState) HandleMove(g *game.Game, p *game.Player, row, col int) error {
	return errors.New("game already has a winner; no more moves allowed")
}
