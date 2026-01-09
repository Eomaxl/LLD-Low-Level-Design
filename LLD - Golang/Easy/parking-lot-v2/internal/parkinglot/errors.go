package parkinglot

import "errors"

var (
	ErrNoSpotAvailable = errors.New("no spot available for vehicle type")
	ErrInvalidTicket   = errors.New("invalid or already closed ticket")
)
