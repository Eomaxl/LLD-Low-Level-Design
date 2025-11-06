package org.solid.srp.strategy;

import org.solid.srp.TaxCalculator;

public class USTaxCalculator implements TaxCalculator {
    public double calculateTax(double amount) {
        return amount * 0.08;
    }
}
