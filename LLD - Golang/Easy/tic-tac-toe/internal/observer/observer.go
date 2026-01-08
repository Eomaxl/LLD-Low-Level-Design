package observer

type GameObserver interface {
	Update(event GameEvent)
}
