package main

import (
	"fmt"
	"tic-tac-toe/internal/game"
	"tic-tac-toe/internal/system"
)

func main() {
	sys := system.GetTicTacToeSystem()

	p1 := &game.Player{Name: "Alice", Symbol: game.X}
	p2 := &game.Player{Name: "Bob", Symbol: game.O}

	sys.CreateGame(p1, p2)

	_ = sys.PrintBoard()
	fmt.Println()

	moves := []struct {
		p    *game.Player
		r, c int
	}{
		{p1, 0, 0},
		{p2, 1, 0},
		{p1, 0, 1},
		{p2, 1, 1},
		{p1, 0, 2}, //Alice wins
	}

	for _, m := range moves {
		if err := sys.MakeMove(m.p, m.r, m.c); err != nil {
			fmt.Println("Move error :", err)
		}
		_ = sys.PrintBoard()
		fmt.Print()
	}

	sys.GetScoreBoard().PrintScores()
}
