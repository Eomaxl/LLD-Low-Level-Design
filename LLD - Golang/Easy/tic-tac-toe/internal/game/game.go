package game

import (
	"errors"
	"fmt"
	"tic-tac-toe/internal/observer"
)

type Game struct {
	board             *Board
	winner            *Player
	state             GameState
	winningStrategies []WinningStrategy

	player1       *Player
	player2       *Player
	status        GameStatus
	currentPlayer *Player

	subject *observer.GameSubject
}

func NewGame(
	p1, p2 *Player,
	boardSize int,
	initialState GameState,
	strategies []WinningStrategy,
	subject *observer.GameSubject,
) *Game {
	return &Game{
		board:             NewBoard(boardSize),
		player1:           p1,
		player2:           p2,
		status:            InProgress,
		currentPlayer:     p1,
		state:             initialState,
		winningStrategies: strategies,
		subject:           subject,
	}
}

func (g *Game) Board() *Board              { return g.board }
func (g *Game) Status() GameStatus         { return g.status }
func (g *Game) Winner() *Player            { return g.winner }
func (g *Game) CurrentPlayer() *Player     { return g.currentPlayer }
func (g *Game) SetState(st GameState)      { g.state = st }
func (g *Game) SetStatus(s GameStatus)     { g.status = s }
func (g *Game) SetWinner(p *Player)        { g.winner = p }
func (g *Game) SwitchPlayer()              { g.switchPlayer() }
func (g *Game) CheckWinner(p *Player) bool { return g.checkWinner(p) }

func (g *Game) MakeMove(player *Player, row, col int) error {
	if g.state == nil {
		return errors.New("game state is nil")
	}
	return g.state.HandleMove(g, player, row, col)
}

func (g *Game) PlaceSymbol(player *Player, row, col int) error {
	if player == nil {
		return errors.New("player is nil")
	}
	return g.board.PlaceSymbol(row, col, player.Symbol)
}

func (g *Game) EnsurePlayersTurn(player *Player) error {
	if g.status != InProgress {
		return fmt.Errorf("game is not in progress (status = %s)", g.status.String())
	}

	if g.currentPlayer != player {
		return fmt.Errorf("not your turn: current player is %s", g.currentPlayer.Name)
	}
	return nil
}

func (g *Game) NotifyFinished() {
	if g.subject == nil {
		return
	}

	ev := observer.GameEvent{
		Status:     g.status.String(),
		WinnerName: "",
	}
	if g.winner != nil {
		ev.WinnerName = g.winner.Name
	}
	g.subject.NotifyObservers(ev)
}

func (g *Game) switchPlayer() {
	if g.currentPlayer == g.player1 {
		g.currentPlayer = g.player2
	} else {
		g.currentPlayer = g.player1
	}
}

func (g *Game) checkWinner(player *Player) bool {
	for _, strat := range g.winningStrategies {
		if strat.CheckWinner(g.board, player) {
			return true
		}
	}
	return false
}
