class Car:
    def __init__(self, brand, model):
        self._brand = brand
        self._model = model
        self._speed = 0

    # Method to acclerate
    def accelerate(self, increment):
        self._speed += increment

    # Method to display info
    def display_status(self):
        print(f"{self._brand} is running at {self._speed} km/h.")
        