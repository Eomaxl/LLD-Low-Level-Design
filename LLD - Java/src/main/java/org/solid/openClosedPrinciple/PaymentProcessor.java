package org.solid.openClosedPrinciple;

public class PaymentProcessor {
    void process(String type, double amount){
        if (type.equals("credit")){
            // credit card logic
        } else if (type.equals("paypal")){
            // paypal logic
        }
        // Adding crypto means modifying the method
    }
}

// Violates Open and Closed principle