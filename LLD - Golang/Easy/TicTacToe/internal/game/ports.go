package game

// GameState = State pattern
type GameState interface {
	HandleMove(g *Game, p *Player, row, col int) error
}

// WinningStrategy = Strategy Pattern
type WinningStrategy interface {
	CheckWinner(board *Board, player *Player) bool
}
