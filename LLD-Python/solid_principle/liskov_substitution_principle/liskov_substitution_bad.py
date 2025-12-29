class Bird:
    def fly(self) -> None:
        # Flying logic
        pass

class Penguin(Bird):
    def fly(self) -> None:
        raise NotImplementedError("Penguins can't fly")