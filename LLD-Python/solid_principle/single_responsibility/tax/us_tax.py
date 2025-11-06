from .tax_calculator import TaxCalculator

class USTaxCalculator(TaxCalculator):
    def calculateTax(selfself, amount:float) -> float:
        return amount * 0.12