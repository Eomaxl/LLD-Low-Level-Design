from abc import ABC, abstractmethod

class Bird(ABC):
    @abstractmethod
    def eat(self) -> None:
        ...
        
class FlyingBird(Bird):
    @abstractmethod
    def fly(self) -> None:
        ...

class Sparrow(FlyingBird):
    def fly(self)-> None:
        pass
    def eat(self)-> None:
        pass

