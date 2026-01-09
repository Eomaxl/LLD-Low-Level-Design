package internal

import "errors"

var (
	ErrNoAvailableSpot = errors.New("no available spot for the vehicle type at gate")
	ErrInvalidTicket   = errors.New("invalid  or already used ticket")
	ErrUnkownGate      = errors.New("unknown gate")
	ErrInvalidConfig   = errors.New("invalid parking lot configuration")
)
