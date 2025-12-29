package liskovsubstitutionprinciple

import "errors"

type Bird struct{}

func (Bird) Fly() error {
	// flying logic
	return nil
}

type PenguinBird struct {
	Bird
}

func (PenguinBird) Fly() error {
	return errors.New("penguin can't fly")
}
