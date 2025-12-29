from abc import ABC, abstractmethod

class PaymentMethod(ABC):
    @abstractmethod
    def process(self, amount: float) -> None:
        ...

class CreditCardPayment(PaymentMethod):
    def process(self, amount: float) -> None:
        # credit card logic
        pass

class PaypalPayment(PaymentMethod):
    def process(self, amount: float) -> None:
        # Paypal logic
        pass

class CryptoPayment(PaymentMethod):
    def process(self, amount: float) -> None:
        # crypto logic
        pass
    
class PaymentProcessor:
    def process(self, amount: float, payment_method: PaymentMethod) -> None:
        payment_method.process(amount)