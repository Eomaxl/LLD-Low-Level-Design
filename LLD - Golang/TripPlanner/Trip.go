package main

type User struct {
	Name string
}

type Driver struct {
	Name string
}

type Trip struct {
	ID              string
	Rider           *User
	Driver          *Driver
	PickUpLocation  *Location
	DropoffLocation *Location
	State           TripState
	Fare            float64
}
