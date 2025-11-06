package org.solid.srp.strategy;

import org.solid.srp.TaxCalculator;

public class IndiaTaxCalculator implements TaxCalculator {
    public double calculateTax(double amount) {
        return amount * 0.18;
    }
}
