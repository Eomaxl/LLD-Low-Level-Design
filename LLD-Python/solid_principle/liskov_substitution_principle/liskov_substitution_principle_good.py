class Workable:
    def work(self) -> None:
        ...

class Feedable:
    def eat(self) -> None:
        ...

class Restable:
    def sleep(self) -> None:
        ...

class Human(Workable, Feedable, Restable):
    def work(self)-> None:
        pass
    def eat(self) -> None:
        pass
    def sleep(self) -> None:
        pass

class Robot(Workable):
    def work(self) -> None:
        pass