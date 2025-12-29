package org.solid.openClosedPrinciple;

public class PaymentProcessorOcp {
    void process(PaymentMethod method, double amount){
        method.process(amount);
    }
}
