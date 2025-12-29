package interfacesegregationprinciple

type Workable interface {
	Work()
}

type Restable interface {
	sleep()
}

type Feedable interface {
	Eat()
}

type Human struct{}

func (Human) Work()     {}
func (Human) Restable() {}
func (Human) Feedable() {}

type Robot struct{}

func (Robot) Work() {}
