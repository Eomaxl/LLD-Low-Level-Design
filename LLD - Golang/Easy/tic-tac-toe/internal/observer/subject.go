package observer

type GameSubject struct {
	observers []GameObserver
}

func NewGameSubject() *GameSubject {
	return &GameSubject{observers: make([]GameObserver, 0)}
}

func (s *GameSubject) AddObserver(o GameObserver) {
	s.observers = append(s.observers, o)
}

func (s *GameSubject) RemoveObserver(o GameObserver) {
	for i := range s.observers {
		if s.observers[i] == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			return
		}
	}
}

func (s *GameSubject) NotifyObservers(event GameEvent) {
	for _, o := range s.observers {
		o.Update(event)
	}
}
