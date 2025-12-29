class PaymentProcessorBad:
    def process(self, amount: float, payment_type: str) -> None:
        if payment_type == "credit card":
            # credit card info
            pass
        elif payment_type == "paypal":
            # paypal info
            pass
        
        