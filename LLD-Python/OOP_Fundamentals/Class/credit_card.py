class CreditCard:
    def __init__(self, cardNumber:int, name: string, cvv: int, date_mon:int):
        self._name = name
        self._cardNumber = cardNumber
        self._cvv = cvv
        self._date_mon = date_mon
        self._total_credit = 10000

    def process_payment(self, amount:int):
        if self._total_credit > amount:
            self._total_credit -= amount
            return True
        else:
            return False