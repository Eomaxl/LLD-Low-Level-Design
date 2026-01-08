package system

import (
	"errors"
	"fmt"
	"sync"
	"tic-tac-toe/internal/game"
	"tic-tac-toe/internal/game/state"
	"tic-tac-toe/internal/game/strategy"
	"tic-tac-toe/internal/observer"
)

type TicTacToeSystem struct {
	scoreboard *observer.Scoreboard
	subject    *observer.GameSubject
	game       *game.Game
}

var (
	instance *TicTacToeSystem
	once     sync.Once
)

func GetTicTacToeSystem() *TicTacToeSystem {
	once.Do(func() {
		sb := observer.NewScoreboard()
		sub := observer.NewGameSubject()
		sub.AddObserver(sb)

		instance = &TicTacToeSystem{
			scoreboard: sb,
			subject:    sub,
		}
	})
	return instance
}

func (s *TicTacToeSystem) GetScoreBoard() *observer.Scoreboard {
	return s.scoreboard
}

func (s *TicTacToeSystem) CreateGame(p1, p2 *game.Player) {
	strats := []game.WinningStrategy{
		strategy.ColumnWinningStrategy{},
		strategy.DiagonalWinningStrategy{},
		strategy.RowWinningStrategy{},
	}

	s.game = game.NewGame(p1, p2, 3, state.InProgressState{}, strats, s.subject)
}

func (s *TicTacToeSystem) MakeMove(player *game.Player, row, col int) error {
	if s.game == nil {
		return errors.New("no active game : call CreateGame first")
	}
	return s.game.MakeMove(player, row, col)
}

func (s *TicTacToeSystem) PrintBoard() error {
	if s.game == nil {
		return errors.New("no active game")
	}

	s.game.Board().PrintBoard()
	fmt.Println("Status:", s.game.Status().String())

	if s.game.Status() == game.InProgress {
		fmt.Println("Turn :", s.game.CurrentPlayer().Name)
	}

	if w := s.game.Winner(); w != nil {
		fmt.Println("Winner :", w.Name)
	}

	return nil
}
