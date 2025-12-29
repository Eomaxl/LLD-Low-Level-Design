package main

import "fmt"

// Car struct (similar to class)
type Car struct {
	brand string // unexpected (private to package)
	model string
	speed int
}

// Constructor-like function
func NewCar(brand, model string) *Car {
	return &Car{
		brand: brand,
		model: model,
		speed: 0,
	}
}

// Method to accelerate
func (c *Car) Accelerate(increment int) {
	c.speed += increment
}

// Method to display status
func (c *Car) DisplayStatus() {
	fmt.Printf("%s is running at %d km/h.\n", c.brand, c.speed)
}

// func main() {
// 	car := NewCar("BMW", "X5")
// 	car.Accelerate(50)
// 	car.DisplayStatus()
// }
