class Worker:
    def work(self) -> None:
        ...
    
    def eat(self) -> None:
        ...

    def sleep(self) -> None:
        ...

class Robot(Worker):
    def work(self) -> None:
        pass

    def eat(self) -> None:
        pass

    def sleep(self) -> None:
        pass