package state

import (
	"errors"
	"tic-tac-toe/internal/game"
)

type DrawState struct{}

func (st DrawState) HandleMove(g *game.Game, p *game.Player, row, col int) error {
	return errors.New("game ended in a draw; no more moves allowed")
}
