from abc import ABC, abstractmethod

class MessageSender(ABC):
    @abstractmethod
    def send(self, message: str) -> None:
        ...

class EmailSender(MessageSender):
    def send(self, message: str) -> None:
        # send email
        pass

class NotificationService:
    def __init__(self, sender: MessageSender) -> None:
        self.sender = sender

    def notify(self, message: str)-> None:
        self.sender.send(message)