from tax.tax_calculator import TaxCalculator

class Invoice:
    def __init__(self, amount: float, taxCalculator: TaxCalculator):
        self.amount = amount
        self.taxCalculator = taxCalculator

    def getTotalAmount(self) -> float:
        return self.amount + self.taxCalculator.calculateTax(self.amount)