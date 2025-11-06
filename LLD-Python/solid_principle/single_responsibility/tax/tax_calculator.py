from abc import ABC, abstractmethod

class TaxCalculator(ABC):
    @abstractmethod
    def calculateTax(selfself, amount:float) -> float:
        pass