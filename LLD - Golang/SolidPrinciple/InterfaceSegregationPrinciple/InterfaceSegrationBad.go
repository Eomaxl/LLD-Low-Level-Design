package interfacesegregationprinciple

type Worker interface {
	Work()
	Eat()
	Sleep()
}

type RobotBad struct{}

func (RobotBad) Work() {}

func (RobotBad) Eat() {}

func (RobotBad) Sleep() {}
