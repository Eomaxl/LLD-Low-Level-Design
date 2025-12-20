from abc import ABC, abstractmethod
import math

class PricingCalculator(ABC):
    @abstractmethod
    def calculate_fare(self, trip: "Trip") -> float:
        pass

class StandardPricingCalculator(PricingCalculator):
    BASE_FARE = 2.50
    COST_PER_MILE = 1.50
    COST_PER_MINUTE = 0.25

    def calculate_fare(self, trip: "Trip") -> float:
        distance = self.calculate_distance(
            trip.pickup_location,
            trip.dropoff_location
        )
        duration = self._estimate_duration(distance)

        return (self.BASE_FARE + 
                (duration * self.COST_PER_MINUTE) + 
                (distance * self.COST_PER_MILE))
    
    def _calculate_distance(self, from_loc: "Location", to_loc: "Location") -> float:
        # Simplified distance calculation
        lat_diff = to_loc.latitude - from_loc.latitude
        lon_diff = to_loc.longitude - from_loc.longitude
        return Math.sqrt(lat_diff ** 2 + lon_diff ** 2) * 69
    
    def _estimate_duration(self, distance_miles: float) -> float:
        return (distance_miles / 30.0) * 60