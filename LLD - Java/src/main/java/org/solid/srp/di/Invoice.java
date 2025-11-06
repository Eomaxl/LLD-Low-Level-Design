package org.solid.srp.di;

import org.solid.srp.TaxCalculator;

public class Invoice {
    private double amount;
    private TaxCalculator  taxCalculator;

    public Invoice(double amount, TaxCalculator taxCalculator) {
        this.amount = amount;
        this.taxCalculator = taxCalculator;
    }

    public double getAmount() {
        return amount + taxCalculator.calculateTax(amount);
    }
}
