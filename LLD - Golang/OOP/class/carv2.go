package main

import (
	"errors"
	"fmt"
)

type Carv2 struct {
	brand string
	model string
	speed int
}

// Constructor like method
func NewCarv2(brand, model string) *Carv2 {
	return &Carv2{
		brand: brand,
		model: model,
		speed: 0,
	}
}

// Setter
func (c *Carv2) SetSpeed(value int) error {
	if value < 0 {
		return errors.New("Speed cannot be negative")
	}
	c.speed = value
	return nil
}

// Getter
func (c *Carv2) Speed() int {
	return c.speed
}

// Business Method
func (c *Carv2) Accelerate(increment int) error {
	return c.SetSpeed(c.speed + increment)
}

func (c *Carv2) DisplayStatus() {
	fmt.Printf("%s is running at %d km/h.\n", c.brand, c.speed)
}

func main() {
	car := Carv2("Suzuki", "Baleno")
	car.SetSpeed(50)
	fmt.Printf(car.Speed())

	err := car.SetSpeed(-10)
	if err != nil {
		fmt.Println("Error : ", err)
	}
}
