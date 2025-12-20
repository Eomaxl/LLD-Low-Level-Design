from enum import Enum
from dataclasses import dataclass

class Trip:
    def __init__(self, id:str, rider: "User", pickup: "Location", dropoff: "Location"):
        self.id = id
        self.rider = rider
        self.driver = None
        self.pickup_location = pickup
        self.dropoff_location = dropoff
        self.state = TripState.REQUESTED
        self.fare = 0.0

    def assign_driver(self, driver: "Driver"):
        if self.state != TripState.REQUESTED :
            raise ValueError(f"Cannot assign driver to trip in state : {self.state}")
        self.driver = driver
        self.state = Trip.Driver_ASSIGNED

    def start_trip(self):
        if self.state != TripState.DRIVER_ASSIGNED:
            raise ValueError(f"Cannot start trip in state: {self.state}")
        self.state = TripState.IN_PROGRESS

    def complete_trip(self, calculator: "PricingCalculator"):
        if self.state != TripState.IN_PROGRESS:
            raise ValueError(f"Cannot complete trip in state: {self.state}")
        self.fare = calculator.calculate_fare(self)
        self.state = TripState.COMPLETED

    def cancel_trip(self):
        if self.state == TripState.COMPLETED:
            raise ValueError("Cannot cancel completed trip")
        self.state = TripState.CANCELLED
        