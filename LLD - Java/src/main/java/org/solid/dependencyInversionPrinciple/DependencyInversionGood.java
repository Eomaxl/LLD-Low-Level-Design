package org.solid.dependencyInversionPrinciple;

public class DependencyInversionGood {
}

interface MessageSender {
    void send(String message);
}

class EmailSender implements MessageSender {
    public void send(String message) { /* send email */}
}

class NotificationService {
    private final MessageSender sender;

    NotificationService(MessageSender sender){
        this.sender = sender;
    }

    void notify(String message){
        sender.send(message);
    }
}

