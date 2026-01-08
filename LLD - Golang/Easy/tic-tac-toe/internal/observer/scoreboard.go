package observer

import "fmt"

type Scoreboard struct {
	scores map[string]int
	draws  int
}

func NewScoreboard() *Scoreboard {
	return &Scoreboard{scores: map[string]int{}}
}

func (s *Scoreboard) Update(event GameEvent) {
	switch event.Status {
	case "WINNER_X", "WINNER_O":
		if event.WinnerName != "" {
			s.scores[event.WinnerName]++
		}
	case "DRAW":
		s.draws++
	}
}

func (s *Scoreboard) PrintScores() {
	fmt.Println("== SCOREBOARD ==")
	for name, score := range s.scores {
		fmt.Printf("%s: %d\n", name, score)
	}
	fmt.Printf("Draws: %d\n", s.draws)
	fmt.Println("================")
}
