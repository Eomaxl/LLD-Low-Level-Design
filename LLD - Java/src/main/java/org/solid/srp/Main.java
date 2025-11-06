package org.solid.srp;

import org.solid.srp.di.Invoice;
import org.solid.srp.strategy.IndiaTaxCalculator;
import org.solid.srp.strategy.UKTaxCalculator;
import org.solid.srp.strategy.USTaxCalculator;

public class Main {
    public static void main(String[] args) {
        double amount = 1000;

        Invoice indiaInvoice = new Invoice(amount, new IndiaTaxCalculator());
        System.out.println("Total India : "+indiaInvoice.getAmount());

        Invoice usInvoice = new Invoice(amount, new USTaxCalculator());
        System.out.println("Total US : "+usInvoice.getAmount());

        Invoice ukInvoice = new Invoice(amount, new UKTaxCalculator());
        System.out.println("Total UK : "+ukInvoice.getAmount());
    }
}
