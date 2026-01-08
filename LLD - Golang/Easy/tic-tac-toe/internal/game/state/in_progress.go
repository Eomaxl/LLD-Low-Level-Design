package state

import "tic-tac-toe/internal/game"

type InProgressState struct{}

func (st InProgressState) HandleMove(g *game.Game, p *game.Player, row, col int) error {
	if err := g.EnsurePlayersTurn(p); err != nil {
		return err
	}

	if err := g.PlaceSymbol(p, row, col); err != nil {
		return err
	}

	if g.CheckWinner(p) {
		g.SetWinner(p)
		g.SetStatus(winnerStatus(p.Symbol))
		g.SetState(WinnerState{})
		g.NotifyFinished()
		return nil
	}

	if g.Board().IsFull() {
		g.SetState(DrawState{})
		g.SetStatus(game.Draw)
		g.NotifyFinished()
		return nil
	}

	g.SwitchPlayer()
	return nil
}

func winnerStatus(sym game.Symbol) game.GameStatus {
	if sym == game.X {
		return game.WinnerX
	}
	return game.WinnerO
}
