class Car:
    def __init__(self,brand, model):
        self._brand = brand
        self._model = model
        self._speed = 0

    # Getter for speed
    @property
    def speed(self):
        return self._speed
    
    # Setter for speed
    @speed.setter
    def speed(self, value):
        if value < 0:
            raise ValueError("Speed cannot be negative")
        self._speed = value
    

    def accelerate(self, increment):
        self.speed += increment

    def display_status(self):
        print(f"{self._brand} is running at {self._speed} km/h.")

def main():
    car = Car("BMW","X5")
    car.speed(50)
    print(car.speed)

    car.speed(-10)