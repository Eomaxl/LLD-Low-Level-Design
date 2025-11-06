from .tax_calculator import TaxCalculator

class IndiaTaxCalculator(TaxCalculator):
    def calculateTax(selfself, amount:float) -> float:
        return amount * 0.18