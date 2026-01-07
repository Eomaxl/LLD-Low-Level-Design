package game

type Symbol int

const (
	Empty Symbol = iota
	X
	O
)

func (s Symbol) Char() rune {
	switch s {
	case X:
		return 'X'
	case O:
		return 'O'
	default:
		return '.'
	}
}

type GameStatus int

const (
	InProgress GameStatus = iota
	WinnerX
	WinnerO
	Draw
)

func (gs GameStatus) String() string {
	switch gs {
	case InProgress:
		return "IN_PROGRESS"
	case WinnerX:
		return "WINNER_X"
	case WinnerO:
		return "WINNER_O"
	case Draw:
		return "DRAW"
	default:
		return "UNKNOWN"
	}
}
