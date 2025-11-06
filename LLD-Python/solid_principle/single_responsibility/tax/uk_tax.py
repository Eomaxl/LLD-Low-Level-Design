from .tax_calculator import TaxCalculator

class UKTaxCalculator(TaxCalculator):
    def calculateTax(selfself, amount:float) -> float:
        return amount * 0.08