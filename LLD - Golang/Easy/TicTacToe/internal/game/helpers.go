package game

func winnerStatusFor(sym Symbol) GameStatus {
	if sym == X {
		return WinnerX
	}
	return WinnerO
}
