from invoice.invoice import Invoice
from tax.india_tax import IndiaTaxCalculator
from tax.uk_tax import UKTaxCalculator
from tax.us_tax import USTaxCalculator

def main():
    amount = 1000

    india_invoice = Invoice(amount, IndiaTaxCalculator())
    print(f"Total (India): ₹{india_invoice.getTotalAmount():.2f}")

    us_invoice = Invoice(amount, USTaxCalculator())
    print(f"Total (US): ${us_invoice.getTotalAmount():.2f}")

    uk_invoice = Invoice(amount, UKTaxCalculator())
    print(f"Total (UK): £{uk_invoice.getTotalAmount():.2f}")

if __name__ == "__main__":
    main()