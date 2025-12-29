package liskovsubstitutionprinciple

type BirdGood interface {
	Eat()
}

type FlyingBird interface {
	BirdGood
	Fly()
}

type Sparrow struct{}

func (Sparrow) Eat() {}
func (Sparrow) Fly() {}

type PenguinGood struct{}

func (PenguinGood) Eat() {}
