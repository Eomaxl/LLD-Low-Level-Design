class EmailSender:
    def send(self, message: str) -> None:
        # send email
        pass

class NotificationService:
    def __init__(self)-> None:
        self.email_sender = EmailSender()

    def notify(self, message: str) -> None:
        self.email_sender.send(message)